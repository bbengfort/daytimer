package daytimer

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

// EditFile creates a temporary file by copying the contents of the
// specified file to a temporary directory it then execs an editor on the
// temp file, and if the file is closed without errors, copies the temporary
// file back to the original location.
//
// Trying to provide similar functionality to crontab -e or git commit.
func EditFile(path string) error {
	// Find the editor to use
	editor, err := findEditor()
	if err != nil {
		return err
	}

	// Create the temporary directory and ensure we clean up when done.
	tmpDir := os.TempDir()
	defer os.RemoveAll(tmpDir)

	// Get the temporary file location
	tmpFile := filepath.Join(tmpDir, filepath.Base(path))

	// Copy the original file to the tmpFile
	if err = copyFile(path, tmpFile); err != nil {
		return err
	}

	// Create the editor command
	cmd := exec.Command(editor, tmpFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the editor command and wait for it to finish
	if err = cmd.Start(); err != nil {
		return err
	}

	if err = cmd.Wait(); err != nil {
		return err
	}

	// Copy the tmp file back to the original file
	return copyFile(tmpFile, path)
}

// List of default editors to search for on the path if $EDITOR or configured
// editor is not specified.
var editors = [4]string{"vim", "emacs", "nano"}

// Looks up the path to the editor to use, returning an error if one cannot
// be found (or something else goes wrong along the way).
func findEditor() (string, error) {

	// First look it up in the configuration
	config, err := LoadConfig()
	if err != nil {
		return "", err
	}

	if config.Editor != "" {
		return config.Editor, nil
	}

	// Next, look up the editor in the environment
	if editor := os.Getenv("EDITOR"); editor != "" {
		return editor, nil
	}

	// Finally go through all of the options above and see if they exist.
	for _, name := range editors {
		path, err := exec.LookPath(name)
		if err == nil {
			return path, nil
		}
	}

	return "", errors.New("no editor found")
}

// Copies a file from source to dest.
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, in); err != nil {
		return err
	}

	return nil
}

package daytimer

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// Returns the path to the configuration directory.
func configDirectory() (string, error) {
	// Get the user to look up the user's home directory
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	// Get the hidden credentials directory, making sure it's created
	confDir := filepath.Join(usr.HomeDir, ".daytimer")
	os.MkdirAll(confDir, 0700)

	return confDir, nil
}

// Returns an iterator that returns one line of a file at a time.
func readLines(path string, skip bool) (<-chan string, error) {
	fobj, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fobj)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	chnl := make(chan string)
	go func() {
		defer fobj.Close()
		defer close(chnl)

		for scanner.Scan() {
			line := scanner.Text()
			if skip && strings.HasPrefix(line, "#") {
				continue
			}
			chnl <- line
		}
	}()

	return chnl, nil
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

// path exists returns true if a file is at the specified location, note that
// this will return true if there is a permissions error even if the file
// does not exist.
func pathExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

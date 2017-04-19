package daytimer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//===========================================================================
// Prompt Input Helper Functions
//===========================================================================

// Prompt collects a string value from the command line, and continues to
// prompt until a response is provided. Returns empty string on errors.
func Prompt(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s ", prompt)

	response, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	response = strings.TrimSpace(response)
	if response == "" {
		return Prompt(prompt)
	}

	return response, nil
}

// PromptStringValue gets a string value from the command line (defaults to
// the default value) if enter is hit or nothing is submitted.
func PromptStringValue(prompt, value string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s [%s]: ", prompt, value)

	response, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	response = strings.TrimSpace(response)
	if response == "" {
		return value, nil
	}

	return response, nil
}

package textinput

import (
	"fmt"
	"io"
	"os"
)

// GetText reads from an input filename or from STDIN if the filename is empty.
func GetText(infile string) (string, error) {

	var reader io.Reader

	// if file flag is provided, try to open that file
	if infile != "" {
		file, err := os.Open(infile)
		if err != nil {
			return "", fmt.Errorf("Error opening file: %v", err)
		}
		defer file.Close()
		reader = file
	} else {
		// get the word list from STDIN
		stat, _ := os.Stdin.Stat()
		// unless the STDIN is not piped data
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			errString := "No input file specified and no data piped to STDIN\n" +
				"Usage: either pipe data to STDIN or use -f flag to specify input file"
			return "", fmt.Errorf(errString)
		}
		reader = os.Stdin
	}

	bytes, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("Error reading input: %v", err)
	}

	return string(bytes), nil
}

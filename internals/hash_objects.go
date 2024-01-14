package internals

import (
	"errors"
	"os"
)

// For hash-object command
// Creates a blob object
func HashObject(file string) (string, error) {
	fileData, err := os.ReadFile(file)
	if err != nil {
		return "", errors.New("Error reading file.\n")
	}
	_, sha := writeObject("blob", fileData)

	return sha, nil
}

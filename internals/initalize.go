package internals

import (
	"fmt"
	"os"
)

// For init command
func InitalizeGit() {
	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		if err := os.Mkdir(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
		}
	}

	HeadFileContents := []byte("ref: refs/heads/master\n")
	if err := os.WriteFile(".git/HEAD", HeadFileContents, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
	}
}

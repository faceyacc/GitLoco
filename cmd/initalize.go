package cmd

import (
	"fmt"
	"os"
)

func initalizeGit() {
	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		if err := os.Mkdir(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
		}
	}
}

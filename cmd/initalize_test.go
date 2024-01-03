package cmd

import (
	"os"
	"testing"
)

func TestInitalizeGit(t *testing.T) {
	t.Run("initalize .git directory", func(t *testing.T) {
		initalizeGit()

		for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				t.Error(".git directory or subdirectories were not created")
			}
			defer os.RemoveAll(dir)
		}

	})
}

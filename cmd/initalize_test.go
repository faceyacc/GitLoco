package cmd

import (
	"os"
	"testing"
)

func TestInitalizeGit(t *testing.T) {
	dirs := []string{".git", ".git/objects", ".git/refs"}
	t.Run("initalize .git directory", func(t *testing.T) {
		initalizeGit()
		for _, dir := range dirs {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				t.Error(".git directory or subdirectories were not created")
			}
			defer os.RemoveAll(dir)
		}
	})

	t.Run("create HEAD file", func(t *testing.T) {
		initalizeGit()
		HEADfile := ".git/HEAD"
		for _, dir := range dirs {
			defer os.RemoveAll(dir)
		}

		if _, err := os.Stat(HEADfile); os.IsNotExist(err) {
			t.Error("HEAD file was not created")
		}
	})
}

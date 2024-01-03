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

func TestCatfile(t *testing.T) {
	test_blob_hash := "3d21ec53a331a6f037a91c368710b99387d012c1"
	// test_string := "dooby scooby vanilla dooby vanilla humpty"

	t.Run("Contruct file path to blob object", func(t *testing.T) {
		got, _ := constructBlob(test_blob_hash)
		expected := ".git/objects/3d/21ec53a331a6f037a91c368710b99387d012c1"

		if got != expected {
			t.Errorf("got %v, but expected %v", got, expected)
		}
	})

}

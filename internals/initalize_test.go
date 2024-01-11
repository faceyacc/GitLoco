package internals

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/Flaque/filet"
)

func TestInitalizeGit(t *testing.T) {
	dirs := []string{".git", ".git/objects", ".git/refs"}
	t.Run("initalize .git directory", func(t *testing.T) {
		InitalizeGit()
		for _, dir := range dirs {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				t.Error(".git directory or subdirectories were not created")
			}
			defer os.RemoveAll(dir)
		}
	})

	t.Run("create HEAD file", func(t *testing.T) {
		InitalizeGit()
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

	t.Run("contruct file path to blob object", func(t *testing.T) {
		blob_filepath, _ := constructObjectsFile(test_blob_hash)
		got := fmt.Sprintf(blob_filepath)
		expected := ".git/objects/3d/21ec53a331a6f037a91c368710b99387d012c1"

		if got != expected {
			t.Errorf("got %v, but expected %v", got, expected)
		}
	})

	t.Run("error if file doesnt exist", func(t *testing.T) {
		assertError := func(t testing.TB, got error, want string) {
			t.Helper()
			if got == nil {
				t.Fatal("didn't get an error but wanted one")
			}

			if got.Error() != want {
				t.Errorf("got %q, want %q", got, want)
			}
		}
		_, got := CatFile(test_blob_hash)

		assertError(t, got, "Error reading file\n")
	})
}

func TestHasObject(t *testing.T) {
	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("read file", func(t *testing.T) {

		_, err := HashObject("expected")

		assertError(t, err, "Error reading file.\n")
	})

	t.Run("added blob header to file", func(t *testing.T) {
		defer filet.CleanUp(t)
		testFile := filet.File(t, "test.txt", "hej hej")
		file_data, _ := os.ReadFile(testFile.Name())

		got := addHeader(file_data)
		expected := fmt.Sprintf("blob %v\x00%v", len(file_data), string(file_data))

		if got != expected {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	})

	t.Run("length of sha-1", func(t *testing.T) {
		defer filet.CleanUp(t)
		testFile := filet.File(t, "test.txt", "hej hej")
		file_data, _ := os.ReadFile(testFile.Name())
		data := addHeader(file_data)

		got := len(calculateSha(data))
		expected := 40

		if got != expected {
			t.Fatalf("Expected sha hash to be %v characters long, got %v", expected, got)
		}
	})

	t.Run("generate directory", func(t *testing.T) {
		// regex to match dir: /.git\/objects\/[0-9]{1}[A-Za-z]{1}\/[0-9a-z]{38}
		defer filet.CleanUp(t)
		testFile := filet.File(t, "test.txt", "hej hej")
		file_data, _ := os.ReadFile(testFile.Name())
		data := addHeader(file_data)
		sha := calculateSha(data)

		test_dir := generateDirFromSha(sha)

		// match test_dir pattern with regex
		matched, _ := regexp.MatchString(`.git\/objects\/[0-9]{1}[A-Za-z]{1}`, test_dir)
		if matched == false {
			t.Errorf("Incorrect directory path. Got %v, but should be in\nthe following format: '.git/objects/0a'", test_dir)
		}
	})
}

func TestWriteTree(t *testing.T) {
}

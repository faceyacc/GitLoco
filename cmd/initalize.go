package cmd

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
)

func initalizeGit() {
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

func constructBlob(blob_hash string) (string, string, error) {
	dir_name := blob_hash[:2]
	file_name := blob_hash[2:]
	blob_filepath := fmt.Sprintf(".git/objects/%v/%v", dir_name, file_name)
	return blob_filepath, file_name, nil
}

func catfile(blob_hash string) (string, error) {

	// Construct the file path to the blob object using the hash
	blob_filepath, file_name, err := constructBlob(blob_hash)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error constructing blob filepath\n")
		return "", err
	}

	// Try to read file
	file, err := fs.ReadFile(os.DirFS(blob_filepath), file_name)
	if err != nil {
		// fmt.Fprintf(os.Stderr, "Error reading file\n")
		return "", errors.New("Error reading file\n")
	}
	fmt.Printf("print test: %v", file)

	// Create a new zlib reader with zlib.NewReader
	zlib_reader, err := zlib.NewReader(bytes.NewReader(file))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decompressing blob for sha:%v\n", blob_hash)
	}

	// Read the decompressed blob content
	decompress, err := io.ReadAll(zlib_reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading blob for sha:%v\n", blob_hash)
	}
	defer zlib_reader.Close()

	// fmt.Print() the content
	object := string(decompress)
	fmt.Print(object)

	return "", nil
}

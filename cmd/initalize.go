package cmd

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
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

func constructBlob(blob_hash string) (string, error) {
	dir_name := blob_hash[:2]
	file_name := blob_hash[2:]
	blob_filepath := fmt.Sprintf(".git/objects/%v/%v", dir_name, file_name)
	return blob_filepath, nil
}

func catfile(blob_hash string) string {

	// Construct the file path to the blob object using the hash
	blob_filepath, err := constructBlob(blob_hash)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error constructing blob filepath\n")
	}

	// Try to read file
	file, err := os.ReadFile(blob_filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file\n")
	}

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

	return ""
}

package cmd

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
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

func constructBlob(blob_hash string) (string, string) {
	dir_name := blob_hash[:2]
	file_name := blob_hash[2:]
	blob_filepath := fmt.Sprintf(".git/objects/%v/%v", dir_name, file_name)
	return blob_filepath, file_name
}

func catfile(blob_hash string) (string, error) {

	// Construct the file path to the blob object using the hash
	blob_filepath, file_name := constructBlob(blob_hash)

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

func addHeader(fileData []byte) string {
	data := fmt.Sprintf("blob %v\x00%v", len(fileData), string(fileData))
	return data
}

func calculateSha(fileData string) string {
	hasher := sha1.New()
	hasher.Write([]byte(fileData))
	hash := hasher.Sum(nil)
	sha := hex.EncodeToString(hash)
	return sha
}

func hashobject(file string) (string, error) {
	// Read in file given (<file>) using os.ReadFile
	fileData, err := os.ReadFile(file)
	if err != nil {
		return "", errors.New("Error reading file.\n")
	}

	// Add the blob header too file data (i.e. "blob"+filesize+"\x00")
	data := addHeader(fileData)

	// Calculate the file's SHA-1 hash using crypto/sha1
	hasher := sha1.New()
	hasher.Write([]byte(data))
	hash := hasher.Sum(nil)
	sha := hex.EncodeToString(hash)

	return string(rune(len(sha))), nil
}

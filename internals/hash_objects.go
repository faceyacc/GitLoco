package internals

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"os"
	"path"
)

// For hash-object command
// Creates a blob object
func HashObject(file string) (string, error) {
	// Read in file given (<file>) using os.ReadFile
	fileData, err := os.ReadFile(file)
	if err != nil {
		return "", errors.New("Error reading file.\n")
	}

	// Add the blob header too file data (i.e. "blob"+filesize+"\x00")
	data := addHeader(fileData)

	// Calculate the file's SHA-1 hash using crypto/sha1
	sha := calculateSha(data)

	// generate the path in ".git/objects" directory
	// based on the first to characters and the rest
	// of the SHA hash (i.e. ".git/objects/3d/21ec53a331a6f037a91c368710b99387d012c1")

	dir := generateDirFromSha(sha)

	// Make directory to using sha
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
	}

	// compress the data using zlib
	var buffer bytes.Buffer
	w := zlib.NewWriter(&buffer)
	w.Write([]byte(data))
	w.Close()

	blobFileName := sha[2:]
	fullPath := path.Join(dir, blobFileName)

	//  write data to file in .git/objects directory
	os.WriteFile(fullPath, buffer.Bytes(), 0755)

	// fmt.Print() the 40-character string from zlibs
	return sha, nil
}

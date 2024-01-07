package internals

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
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

func constructBlob(blob_hash string) (string, string) {
	dir_name := blob_hash[:2]
	file_name := blob_hash[2:]
	blob_filepath := fmt.Sprintf(".git/objects/%v/%v", dir_name, file_name)
	return blob_filepath, file_name
}

// For cat-file command
// Reads a blob object
func Catfile(blob_hash string) (string, error) {

	// Construct the file path to the blob object using the hash
	blob_filepath, _ := constructBlob(blob_hash)

	// Try to read file
	file, err := os.ReadFile(blob_filepath)
	if err != nil {
		// fmt.Fprintf(os.Stderr, "Error reading file\n")
		return "", errors.New("Error reading file\n")
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
	split_res := bytes.Split(decompress, []byte("\x00"))
	object := string(split_res[len(split_res)-1])

	return object, nil
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

func generateDirFromSha(sha string) string {
	blobDir := sha[:2]

	dirName := fmt.Sprintf(".git/objects/%v", blobDir)

	return dirName
}

// For hash-object command
// Creates a blob object
func Hashobject(file string) (string, error) {
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

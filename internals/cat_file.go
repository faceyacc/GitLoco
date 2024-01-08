package internals

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"os"
)

// For cat-file command.
// Reads a blob object
func CatFile(blob_hash string) (string, error) {

	// Construct the file path to the blob object using the hash
	blob_filepath, _ := constructObjectsFile(blob_hash)

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

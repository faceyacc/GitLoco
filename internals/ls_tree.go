package internals

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

// LsTre is used to inspect a tree object
func LsTree(tree_sha string) string {

	// Locate from tree_sha

	treeFilePath, _ := constructObjectsFile(tree_sha)

	// read tree object from tree sha in .git/objects
	file, err := os.ReadFile(treeFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file")
	}

	// Decompress file
	zlibReader, err := zlib.NewReader(bytes.NewBuffer(file))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decompressing blob for sha:%v", tree_sha)
	}

	// Read data
	data, err := io.ReadAll(zlibReader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading blob for sha:%v", tree_sha)
	}
	defer zlibReader.Close()

	// Return Tree entries
	tree_entries := extractTreeEntries(data)
	for _, item := range tree_entries {

		os.Stdout.WriteString(fmt.Sprintf("%s\n", (string(item.Name))))
	}

	return string(data)
}

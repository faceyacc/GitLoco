package internals

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func WriteTree(path string) ([20]byte, string) {
	// Check if in working directory
	// if _, err := os.Stat(".git"); err != nil {
	// 	if os.IsNotExist(err) {
	// 		fmt.Fprintf(os.Stderr, "No .git repository initalize in this directory.")
	// 	} else {
	// 		fmt.Fprintf(os.Stderr, "Error reading directory: %v", err)
	// 	}
	// }

	treeEntries := []string{}

	dirInfo, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory: %s\n", err)
	}

	for _, item := range dirInfo {
		if item.Name() == ".git" {
			continue
		}

		if item.IsDir() {
			filePath := filepath.Join(path, item.Name())
			hash, _ := WriteTree(filePath)
			row := fmt.Sprintf("%v %v\x00%s", FILE, item.Name(), hash)
			treeEntries = append(treeEntries, row)
		} else {
			filePath := filepath.Join(path, item.Name())
			file, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Fprint(os.Stderr, "Error reading file")
			}
			hash, _ := writeObject("blob", file)
			row := fmt.Sprintf("%v %v\x00%s", FILE, item.Name(), hash)
			treeEntries = append(treeEntries, row)

		}
	}

	// sort.Slice(treeEntries, func(i, j int) bool {
	// 	return treeEntries[i][strings.IndexByte(treeEntries[i], ' ')+1:] < treeEntries[j][strings.IndexByte(treeEntries[i], ' ')+1:]
	// })

	var buffer bytes.Buffer
	for _, entry := range treeEntries {
		buffer.WriteString(entry)
	}
	return writeObject("tree", buffer.Bytes())
}

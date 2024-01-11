package internals

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type TreeEntry struct {
	Mode []byte
	Name []byte
	SHA  []byte
}

const FILE = "100644"
const SUBDIR = "40000"

func getFormatEntry(item []byte) ([]byte, []byte) {
	mode, name, _ := bytes.Cut(item, []byte(" "))
	var modeResult, nameResult string

	// Check if entry has tree header
	if string(mode) == "tree" {
		modeResult = string(mode)
		nameResult = ""
	}

	// Check mode
	if strings.Contains(string(mode), FILE) {
		modeResult = FILE
	}

	if strings.Contains(string(mode), SUBDIR) {
		modeResult = SUBDIR
	}

	// Check name
	isNumerical, _ := regexp.MatchString("[0-9]", string(name))
	if isNumerical == true {
		if strings.Contains(string(name), FILE) {
			modeResult = FILE
		} else if strings.Contains(string(name), SUBDIR) {
			modeResult = SUBDIR
		}
		name_split := strings.Split(string(name), " ")
		nameResult = name_split[len(name_split)-1]

	} else {
		nameResult = string(name)
	}

	return []byte(modeResult), []byte(nameResult)
}

func extractTreeEntries(treeBlob []byte) []TreeEntry {
	var entries []TreeEntry

	items := bytes.Split(treeBlob, []byte("\x00"))

	for _, item := range items {
		getFormatEntry(item)

		if len(item) >= 40 {
			mode, name := getFormatEntry(item)

			entry := TreeEntry{
				Mode: mode,
				Name: name,
				SHA:  item[len(item)-20:],
			}

			entries = append(entries, entry)
		} else {
			mode, name := getFormatEntry(item)

			if string(mode) == "tree" {
				continue
			}
			if len(string(name)) == 0 {
				continue
			}

			dirHeader := []byte("/")
			entry := TreeEntry{
				Mode: mode,
				Name: append(name, dirHeader...),
				SHA:  item[:],
			}
			entries = append(entries, entry)
		}
	}
	return entries
}

func constructObjectsFile(blob_hash string) (string, string) {
	dir_name := blob_hash[:2]
	file_name := blob_hash[2:]
	blob_filepath := fmt.Sprintf(".git/objects/%v/%v", dir_name, file_name)
	return blob_filepath, file_name
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

func writeObject(header string, data []byte) ([20]byte, string) {

	headerBlob := fmt.Sprintf("%s %d\x00", header, len(data))
	blob := append([]byte(headerBlob), data...)

	// Get Blob checksum
	hash := sha1.Sum(blob)

	// Get 20-haracter sha-1
	sha1 := hex.EncodeToString(hash[:])

	if len(sha1) != 40 {
		fmt.Fprintf(os.Stderr, "invalid hash length: %v\n", len(sha1))
		os.Exit(1)
	}

	// Construct and create blob pathusing sha-1
	dir := fmt.Sprintf(".git/objects/%v", sha1[:2])
	filePath := fmt.Sprintf("%v/%v", dir, sha1[2:])

	if err := os.MkdirAll(string(dir), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating dir: %v. got err=%v\n", string(dir), err)
		os.Exit(1)
	}

	// Compress file
	var buf bytes.Buffer
	compressWriter := zlib.NewWriter(&buf)
	_, err := compressWriter.Write(blob)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to buffer: %v\n", err)
		os.Exit(1)
	}
	compressWriter.Close()

	err = os.WriteFile(filePath, buf.Bytes(), 0755)
	if err != nil {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file=%s with error: %v\n", filePath, err)
			os.Exit(1)
		}
	}
	return hash, sha1

}

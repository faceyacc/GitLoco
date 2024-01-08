package internals

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

type TreeEntry struct {
	Mode []byte
	Name []byte
	SHA  []byte
}

const FILE = "100644"
const BLOB = "40000"

func getFormatEntry(item []byte) ([]byte, []byte) {
	mode, name, _ := bytes.Cut(item, []byte(" "))
	var modeResult, nameResult string

	// Chec if entry has tree header
	if string(mode) == "tree" {
		modeResult = string(mode)
		nameResult = ""
	}

	// Check mode
	if strings.Contains(string(mode), FILE) {
		modeResult = FILE
	}

	if strings.Contains(string(mode), BLOB) {
		modeResult = BLOB
	}

	// Check name
	isNumerical, _ := regexp.MatchString("[0-9]", string(name))
	if isNumerical == true {
		if strings.Contains(string(name), FILE) {
			modeResult = FILE
		} else if strings.Contains(string(name), BLOB) {
			modeResult = BLOB
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

		if len(item) > 20 {
			mode, name := getFormatEntry(item)

			entry := TreeEntry{
				Mode: mode,
				Name: name,
				SHA:  item[len(item)-20:],
			}

			entries = append(entries, entry)
		} else if len(item) < 20 {
			mode, name := getFormatEntry(item)
			if string(mode) == "tree" {
				continue
			}
			if len(string(name)) == 0 {
				continue
			}
			entry := TreeEntry{
				Mode: mode,
				Name: name,
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

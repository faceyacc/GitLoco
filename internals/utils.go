package internals

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func constructBlob(blob_hash string) (string, string) {
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

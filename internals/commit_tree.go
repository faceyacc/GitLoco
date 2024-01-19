package internals

import (
	"time"
)

type Commit struct {
	author  string
	email   string
	time    time.Time
	tree    string
	parent  string
	message string
}

func CommitTree(tree_sha string, parent_hash string, message string) string {

	return "dummy sha here"
}

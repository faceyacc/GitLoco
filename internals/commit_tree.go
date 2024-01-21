package internals

import (
	"fmt"
	"strings"
	"time"
)

type commit struct {
	author  string
	email   string
	time    time.Time
	tree    string
	parent  string
	message string
}

func CommitTree(tree_sha, parent_hash, message, author_name, email string) string {

	commit := commit{
		author:  author_name,
		email:   email,
		time:    time.Now(),
		tree:    tree_sha,
		parent:  parent_hash,
		message: message,
	}

	var commitBuilder strings.Builder
	commitBuilder.WriteString(fmt.Sprintf("tree %v\n", commit.author))
	commitBuilder.WriteString(fmt.Sprintf("author %v <%v>\n", commit.author, commit.email))

	// Add optional parent_hash
	if parent_hash != "" {
		commitBuilder.WriteString(fmt.Sprintf("parent %v\n", commit.parent))
	}
	commitBuilder.WriteString(fmt.Sprintf("commiter %v %v\n", commit.author, commit.email))
	commitBuilder.WriteString(fmt.Sprintf("%v\n", commit.time))
	commitBuilder.WriteString(fmt.Sprintf("\n%v\n", commit.message))

	_, sha1 := writeObject("commit", []byte(commitBuilder.String()))

	return sha1
}

package internals

import "time"

type Commit struct {
	author  string
	email   string
	time    time.Time
	tree    string
	parent  string
	message string
}

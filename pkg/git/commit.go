package git

import (
	"fmt"
	"strings"
)

type Commit struct {
	tree      GitSha
	parent    GitSha
	author    string
	committer string
	message   string
}

func GetCommit(base string, sha GitSha) (*Commit, error) {
	object, err := getObject(base, sha)
	if err != nil {
		return nil, err
	}
	commit, err := parseCommit(object)
	if err != nil {
		return nil, err
	}
	return commit, nil
}

func parseCommit(object string) (*Commit, error) {
	treeIndex := strings.Index(object, "\x00tree ")
	parentIndex := strings.Index(object, "\nparent ")
	authorIndex := strings.Index(object, "\nauthor ")
	committerIndex := strings.Index(object, "\ncommitter ")
	messageIndex := strings.Index(object, "\n\n")

	fmt.Printf("treeIndex: %d, parentIndex: %d, authorIndex: %d, committerIndex: %d, messageIndex: %d", treeIndex, parentIndex, authorIndex, committerIndex, messageIndex)

	if treeIndex == -1 || parentIndex == -1 || authorIndex == -1 || committerIndex == -1 || messageIndex == -1 {
		return nil, fmt.Errorf("failed to split commit string:\n%s", object)
	}

	commit := Commit{
		GitSha(object[treeIndex+6 : parentIndex]),
		GitSha(object[parentIndex+8 : authorIndex]),
		object[authorIndex+8 : committerIndex],
		object[committerIndex+11 : messageIndex],
		object[messageIndex+2:],
	}

	return &commit, nil
}

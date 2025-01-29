package git

import (
	"fmt"
	"strings"
)

type Commit struct {
	sha       GitSha
	tree      GitSha
	parent    GitSha
	author    string
	committer string
	message   string
}

// func GetCommit(base: string, sha GitSha) (Commit, error) {
//
// }

func parseCommit(sha GitSha, stringCommit string) (*Commit, error) {
	treeIndex := strings.Index(stringCommit, "\x00tree ")
	parentIndex := strings.Index(stringCommit, "\nparent ")
	authorIndex := strings.Index(stringCommit, "\nauthor ")
	committerIndex := strings.Index(stringCommit, "\ncommitter ")
	messageIndex := strings.Index(stringCommit, "\n\n")

	fmt.Printf("treeIndex: %d, parentIndex: %d, authorIndex: %d, committerIndex: %d, messageIndex: %d", treeIndex, parentIndex, authorIndex, committerIndex, messageIndex)

	if treeIndex == -1 || parentIndex == -1 || authorIndex == -1 || committerIndex == -1 || messageIndex == -1 {
		return nil, fmt.Errorf("failed to split commit string:\n%s", stringCommit)
	}

	commit := Commit{
		sha,
		GitSha(stringCommit[treeIndex+6 : parentIndex]),
		GitSha(stringCommit[parentIndex+8 : authorIndex]),
		stringCommit[authorIndex+8 : committerIndex],
		stringCommit[committerIndex+11 : messageIndex],
		stringCommit[messageIndex+2:],
	}

	return &commit, nil
}

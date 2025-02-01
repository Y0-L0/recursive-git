package git

import (
	"fmt"
	"log/slog"
	"strings"
)

type Commit struct {
	tree        GitSha
	parent      GitSha
	mergeParent GitSha
	author      string
	committer   string
	message     string
}

func newCommit(object string) (*Commit, error) {
	treeIndex := strings.Index(object, "\x00tree ")
	parentIndex := strings.Index(object, "\nparent ")
	authorIndex := strings.Index(object, "\nauthor ")
	committerIndex := strings.Index(object, "\ncommitter ")
	messageIndex := strings.Index(object, "\n\n")

	slog.Debug("Commit string parsing indexes", "treeIndex", treeIndex, "parentIndex", parentIndex, "authorIndex", authorIndex, "committerIndex", committerIndex, "messageIndex", messageIndex)

	if treeIndex == -1 || parentIndex == -1 || authorIndex == -1 || committerIndex == -1 || messageIndex == -1 {
		slog.Warn("Commit message parsing failed", "commit", object)
		return nil, fmt.Errorf("failed to split commit string:\n%s", object)
	}

	parentIndex = parentIndex + 8
	parent := object[parentIndex : parentIndex+40]
	mergeParent := ""
	parentIndex = parentIndex + 48
	if object[parentIndex-8:parentIndex] == "\nparent " {
		mergeParent = object[parentIndex : parentIndex+40]
	}
	parentIndex = parentIndex + 40
	if object[parentIndex:parentIndex+8] == "\nparent " {
		panic("Multiple merge parents are not supported at this stage")
	}

	commit := Commit{
		GitSha(object[treeIndex+6 : treeIndex+6+40]),
		GitSha(parent),
		GitSha(mergeParent),
		object[authorIndex+8 : committerIndex],
		object[committerIndex+11 : messageIndex],
		object[messageIndex+2:],
	}

	slog.Debug("Parsed commit", "commit", commit)
	return &commit, nil
}

package git

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

type Commit struct {
	tree           GitSha
	parent         GitSha
	mergeParent    GitSha
	author         string
	committer      string
	committerEpoch int
	message        string
}

func newCommit(object string) (*Commit, error) {
	treeIndex := strings.Index(object, "\x00tree ")
	parentIndex := strings.Index(object, "\nparent ")
	authorIndex := strings.Index(object, "\nauthor ")
	committerIndex := strings.Index(object, "\ncommitter ")
	messageIndex := strings.Index(object, "\n\n")

	slog.Debug("Commit string parsing indexes", "treeIndex", treeIndex, "parentIndex", parentIndex, "authorIndex", authorIndex, "committerIndex", committerIndex, "messageIndex", messageIndex)

	committer := object[committerIndex+11 : messageIndex]
	epochEnd := strings.LastIndex(committer, " ")
	epochStart := strings.LastIndex(committer[:epochEnd], " ") + 1
	committerEpoch, err := strconv.Atoi(committer[epochStart:epochEnd])
	if err != nil {
		return nil, err
	}

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
		tree:           GitSha(object[treeIndex+6 : treeIndex+6+40]),
		parent:         GitSha(parent),
		mergeParent:    GitSha(mergeParent),
		author:         object[authorIndex+8 : committerIndex],
		committer:      committer,
		committerEpoch: committerEpoch,
		message:        object[messageIndex+2:],
	}

	slog.Debug("Parsed commit", "commit", commit)
	return &commit, nil
}

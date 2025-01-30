package git

import (
	"fmt"
	"log/slog"
	"strings"
)

type Commit struct {
	tree      GitSha
	parent    GitSha
	author    string
	committer string
	message   string
}

func newCommit(object string) (*Commit, error) {
	treeIndex := strings.Index(object, "\x00tree ")
	parentIndex := strings.Index(object, "\nparent ")
	authorIndex := strings.Index(object, "\nauthor ")
	committerIndex := strings.Index(object, "\ncommitter ")
	messageIndex := strings.Index(object, "\n\n")

	slog.Debug("Commit string parsing indexes", "treeIndex", treeIndex, "parentIndex", parentIndex, "authorIndex", authorIndex, "committerIndex", committerIndex, "messageIndex", messageIndex)

	var parents []string
	var parent string
	substring := object[parentIndex:]

	for substring[:8] == "\nparent " {
		parent = substring[8:48]
		parents = append(parents, parent)
		substring = substring[48:]
	}

	slog.Debug("Commit string parsing results", "parents", parents, "substring", substring)

	if treeIndex == -1 || parentIndex == -1 || authorIndex == -1 || committerIndex == -1 || messageIndex == -1 {
		slog.Warn("Commit message parsing failed", "commit", object)
		return nil, fmt.Errorf("failed to split commit string:\n%s", object)
	}

	commit := Commit{
		GitSha(object[treeIndex+6 : treeIndex+6+40]),
		GitSha(parent),
		object[authorIndex+8 : committerIndex],
		object[committerIndex+11 : messageIndex],
		object[messageIndex+2:],
	}

	slog.Debug("Parsed commit", "commit", commit)
	return &commit, nil
}

func (repo *Repo) Commit(sha GitSha) (*Commit, error) {
	object, err := getObject(repo.base, sha)
	if err != nil {
		return nil, err
	}
	commit, err := newCommit(object)
	if err != nil {
		return nil, err
	}
	return commit, nil
}

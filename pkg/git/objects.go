package git

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path"
	"strings"
)

func (repo Repo) Commit(sha GitSha) (*Commit, error) {
	object, err := getObject(repo.base, sha)
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

func getObject(base string, sha GitSha) (string, error) {
	objectPath := path.Join(base, ".git/objects", string(sha[:2]), string(sha[2:]))
	objectBytes, err := os.ReadFile(objectPath)
	if err != nil {
		slog.Error("Failed to read git object file", "sha", sha)
		return "", err
	}

	reader, err := zlib.NewReader(bytes.NewReader(objectBytes))
	if err != nil {
		slog.Error("Failed to initialize zlib reader", "sha", sha)
		return "", err
	}
	defer reader.Close()

	var byteObject bytes.Buffer
	_, err = io.Copy(&byteObject, reader)
	if err != nil {
		slog.Error("Failed to decompress git object", "sha", sha)
		return "", err
	}

	object := byteObject.String()

	slog.Debug("Read git object from file", "sha", sha, "object", object)
	return object, nil
}

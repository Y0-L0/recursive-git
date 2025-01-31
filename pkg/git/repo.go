package git

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type GitSha string

type Repo struct {
	commitCache map[GitSha]*Commit
	branchCache map[string]*Branch
	base        string
	head        GitSha
}

func NewRepo(base string) *Repo {
	return &Repo{
		commitCache: make(map[GitSha]*Commit),
		branchCache: make(map[string]*Branch),
		base:        base,
		head:        GitSha(""),
	}
}

func (repo *Repo) Head() (GitSha, error) {
	if repo.head != GitSha("") {
		return repo.head, nil
	}

	ref, err := os.ReadFile(repo.base + ".git/HEAD")
	if err != nil {
		return "", fmt.Errorf("'base' is not a valid git base directory, %w", err)
	}
	refCleaned := string(ref[5 : len(ref)-1])

	sha, err := repo.ref(refCleaned)
	if err != nil {
		return "", err
	}
	repo.head = sha
	return sha, nil
}

func (repo *Repo) Commit(sha GitSha) (*Commit, error) {
	commit := repo.commitCache[sha]
	if commit != nil {
		slog.Debug("Found commit in commit cache", "sha", sha)
		return commit, nil
	}
	object, err := getObject(repo.base, sha)
	if err != nil {
		return nil, err
	}
	commit, err = newCommit(object)
	if err != nil {
		return nil, err
	}
	repo.commitCache[sha] = commit
	return commit, nil
}

func (repo *Repo) Branch(name string) (*Branch, error) {
	sha, err := repo.ref("refs/heads/" + name)
	if err != nil {
		return nil, err
	}
	branch := newBranch(repo, sha)
	repo.branchCache[name] = branch
	return branch, nil
}

func (repo *Repo) ref(ref string) (GitSha, error) {
	refPath := repo.base + ".git/" + ref
	shaBytes, err := os.ReadFile(refPath)
	if err != nil {
		return "", fmt.Errorf("invalid git ref path, %w", err)
	}
	sha := GitSha(strings.TrimSpace(string(shaBytes)))
	if len(sha) != 40 {
		return "", fmt.Errorf("invalid content of git ref %s", refPath)
	}
	return sha, nil
}

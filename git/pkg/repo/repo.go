package repo

import (
	"fmt"
	"os"
	"strings"
  "github.com/Y0-L0/recursive-git/git/pkg/commit"
  m "github.com/Y0-L0/recursive-git/git/pkg/models"
)

type Repo struct {
	commitCache map[m.GitSha]commit.Commit
	base string
	head m.GitSha
}

func NewRepo(base string) *Repo {
	return &Repo{
		base: base,
		head: m.GitSha(""),
	}
}

func (repo *Repo) Head() (m.GitSha, error) {
	if repo.head != m.GitSha("") {
		return repo.head, nil
	}

	ref, err := os.ReadFile(repo.base + ".git/HEAD")
	if err != nil {
		return "", fmt.Errorf("'base' is not a valid git base directory, %w", err)
	}
	refCleaned := string(ref[5 : len(ref)-1])

	sha, err := repo.ref(refCleaned)
	if err != nil {
		return "wtf: ", err
	}
	repo.head = sha
	return sha, nil
}

func (repo *Repo) ref(ref string) (m.GitSha, error) {
	refPath := repo.base + ".git/" + ref
	shaBytes, err := os.ReadFile(refPath)
	if err != nil {
		return "", fmt.Errorf("invalid git ref path, %w", err)
	}
	sha := m.GitSha(strings.TrimSpace(string(shaBytes)))
	if len(sha) != 40 {
		return "", fmt.Errorf("invalid content of git ref %s", refPath)
	}
	return sha, nil
}

func (repo *Repo) Commit(sha m.GitSha) (*commit.Commit, error) {
	object, err := getObject(repo.base, sha)
	if err != nil {
		return nil, err
	}
	commit, err := commit.NewCommit(object)
	if err != nil {
		return nil, err
	}
	return commit, nil
}

package git

import (
	"fmt"
	"os"
	"strings"
)

type Repo struct {
	commitCache map[GitSha]Commit
	base string
	head GitSha
}

func NewRepo(base string) *Repo {
	return &Repo{
		base: base,
		head: GitSha(""),
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
		return "wtf: ", err
	}
	repo.head = sha
	return sha, nil
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

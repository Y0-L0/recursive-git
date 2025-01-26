package git

import (
	"fmt"
	"os"
	"strings"
)

type GitSha string

func GetHead(base string) (GitSha, error) {
	ref, err := os.ReadFile(base + ".git/HEAD")
	if err != nil {
		return "", fmt.Errorf("'base' is not a valid git base directory, %w", err)
	}
	sha, err := GetRef(base, strings.TrimPrefix(string(ref), "ref: "))
	if err != nil {
		return "", err
	}
	return sha, nil
}

func GetRef(base string, ref string) (GitSha, error) {
	refPath := base + ".git/" + ref
	shaBytes, err := os.ReadFile(refPath)
	if err != nil {
		return "", fmt.Errorf("invalid git ref path %s, %w", refPath, err)
	}
	sha := GitSha(strings.TrimSpace(string(shaBytes)))
	if len(sha) != 40 {
		return "", fmt.Errorf("invalid content of git ref %s", refPath)
	}
	return sha, nil
}

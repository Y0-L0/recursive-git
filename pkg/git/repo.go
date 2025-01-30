package git

import (
	"fmt"
	"os"
	"strings"
)

func GetHead(base string) (GitSha, error) {
	ref, err := os.ReadFile(base + ".git/HEAD")
	if err != nil {
		return "", fmt.Errorf("'base' is not a valid git base directory, %w", err)
	}
	refCleaned := string(ref[5 : len(ref)-1])

	sha, err := getRef(base, refCleaned)
	if err != nil {
		return "wtf: ", err
	}
	return sha, nil
}

func getRef(base string, ref string) (GitSha, error) {
	refPath := base + ".git/" + ref
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

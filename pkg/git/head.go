package git

import (
	"fmt"
	"os"
)

func GetHead(base string) (GitSha, error) {
	ref, err := os.ReadFile(base + ".git/HEAD")
	if err != nil {
		return "", fmt.Errorf("'base' is not a valid git base directory, %w", err)
	}
	refCleaned := string(ref[5:len(ref)-1])

	sha, err := getRef(base, refCleaned)
	if err != nil {
		return "wtf: ", err
	}
	return sha, nil
}

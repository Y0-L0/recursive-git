package git

import (
	"github.com/Y0-L0/recursive-git/testutils"
	"testing"
)

const TEST_REPO_BASE = "./testdata/directory-importer/"

func TestGetRefSuccess(t *testing.T) {
	sha, err := getRef(TEST_REPO_BASE, "refs/heads/main")
	testutils.Ok(t, err)

	testutils.Equals(t, GitSha("21fcd46063d09b0e178619c37bf396beece3a8e2"), sha)
}

func TestGetRefNoFile(t *testing.T) {
	sha, err := getRef(TEST_REPO_BASE, "refs/invalid-location-heads/main")
	testutils.Equals(t, GitSha(""), sha)

	// TODO: Validate that the correct error is returned!
	testutils.Assert(t, err != nil, "Expeced a file not found error but got nil")
}

func TestGetRefWrongContent(t *testing.T) {
	sha, err := getRef(TEST_REPO_BASE, "HEAD")
	testutils.Equals(t, GitSha(""), sha)

	// TODO: Validate that the correct error is returned!
	testutils.Assert(t, err != nil, "Expected an InvalidContent error but got nil")
}


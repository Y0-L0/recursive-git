package git

import (
	"github.com/Y0-L0/recursive-git/testutils"
	"testing"
)

func testRepo() Repo {
	return NewRepo("./testdata/directory-importer/")
}

func TestGetHeadSuccess(t *testing.T) {
	sha, err := testRepo().Head()
	testutils.Ok(t, err)

	testutils.Equals(t, GitSha("21fcd46063d09b0e178619c37bf396beece3a8e2"), sha)
}

func TestGetHeadNoFile(t *testing.T) {
	sha, err := testRepo().ref("refs/invalid-location-heads/main")
	testutils.Equals(t, GitSha(""), sha)

	// TODO: Validate that the correct error is returned!
	testutils.Assert(t, err != nil, "Expeced a file not found error but got nil")
}

func TestGetRefSuccess(t *testing.T) {
	sha, err := testRepo().ref("refs/heads/main")
	testutils.Ok(t, err)

	testutils.Equals(t, GitSha("21fcd46063d09b0e178619c37bf396beece3a8e2"), sha)
}

func TestGetRefNoFile(t *testing.T) {
	sha, err := testRepo().ref("refs/invalid-location-heads/main")
	testutils.Equals(t, GitSha(""), sha)

	// TODO: Validate that the correct error is returned!
	testutils.Assert(t, err != nil, "Expeced a file not found error but got nil")
}

func TestGetRefWrongContent(t *testing.T) {
	sha, err := testRepo().ref("HEAD")
	testutils.Equals(t, GitSha(""), sha)

	// TODO: Validate that the correct error is returned!
	testutils.Assert(t, err != nil, "Expected an InvalidContent error but got nil")
}

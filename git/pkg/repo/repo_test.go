package repo

import (
	"testing"

	m "github.com/Y0-L0/recursive-git/git/pkg/models"
	"github.com/Y0-L0/recursive-git/testutils"
)

func testRepo() *Repo {
	return NewRepo("../../test/testdata/directory-importer/")
}

func TestGetHeadSuccess(t *testing.T) {
	sha, err := testRepo().Head()
	testutils.Ok(t, err)

	testutils.Equals(t, m.GitSha("21fcd46063d09b0e178619c37bf396beece3a8e2"), sha)
}

func TestGetHeadNoFile(t *testing.T) {
	sha, err := testRepo().ref("refs/invalid-location-heads/main")
	testutils.Equals(t, m.GitSha(""), sha)

	// TODO: Validate that the correct error is returned!
	testutils.Assert(t, err != nil, "Expeced a file not found error but got nil")
}

func TestGetRefSuccess(t *testing.T) {
	sha, err := testRepo().ref("refs/heads/main")
	testutils.Ok(t, err)

	testutils.Equals(t, m.GitSha("21fcd46063d09b0e178619c37bf396beece3a8e2"), sha)
}

func TestGetRefNoFile(t *testing.T) {
	sha, err := testRepo().ref("refs/invalid-location-heads/main")
	testutils.Equals(t, m.GitSha(""), sha)

	// TODO: Validate that the correct error is returned!
	testutils.Assert(t, err != nil, "Expeced a file not found error but got nil")
}

func TestGetRefWrongContent(t *testing.T) {
	sha, err := testRepo().ref("HEAD")
	testutils.Equals(t, m.GitSha(""), sha)

	// TODO: Validate that the correct error is returned!
	testutils.Assert(t, err != nil, "Expected an InvalidContent error but got nil")
}

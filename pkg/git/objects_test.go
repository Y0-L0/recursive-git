package git

import (
	"testing"

	"github.com/Y0-L0/recursive-git/testutils"
)

var EXAMPLE_COMMIT = struct {
	sha    GitSha
	object string
	commit Commit
}{
	sha:    GitSha("eeea1494ed65e09bb20d43bd3fc384a3e65da43a"),
	object: "commit 324\x00tree c9be4a262387887608b3c3c3d1237b43f5d3ac82\nparent 2c6bd14b0015249b232685b50ab69016e74cc775\nauthor Johannes Lohmer <lohmer@univention.de> 1736442888 +0100\ncommitter Johannes Lohmer <lohmer@univention.de> 1736443347 +0100\n\ntest: Create and delete the example.org maildomain as part of every testrun with an autouse fixture\n",
	commit: Commit{
		GitSha("c9be4a262387887608b3c3c3d1237b43f5d3ac82"),
		GitSha("2c6bd14b0015249b232685b50ab69016e74cc775"),
		"Johannes Lohmer <lohmer@univention.de> 1736442888 +0100",
		"Johannes Lohmer <lohmer@univention.de> 1736443347 +0100",
		"test: Create and delete the example.org maildomain as part of every testrun with an autouse fixture\n",
	},
}

func TestParseCommit(t *testing.T) {

	commit, err := parseCommit(EXAMPLE_COMMIT.object)
	testutils.Ok(t, err)

	testutils.Equals(t, &EXAMPLE_COMMIT.commit, commit)
}

func TestGetObject(t *testing.T) {
	expected := "commit 324\x00tree c9be4a262387887608b3c3c3d1237b43f5d3ac82\nparent 2c6bd14b0015249b232685b50ab69016e74cc775\nauthor Johannes Lohmer <lohmer@univention.de> 1736442888 +0100\ncommitter Johannes Lohmer <lohmer@univention.de> 1736443347 +0100\n\ntest: Create and delete the example.org maildomain as part of every testrun with an autouse fixture\n"

	obj, err := getObject(TEST_REPO_BASE, GitSha("eeea1494ed65e09bb20d43bd3fc384a3e65da43a"))
	testutils.Ok(t, err)

	testutils.Equals(t, expected, obj)
}

func TestGetCommit(t *testing.T) {
	commit, err := GetCommit(TEST_REPO_BASE, EXAMPLE_COMMIT.sha)
	testutils.Ok(t, err)

	testutils.Equals(t, &EXAMPLE_COMMIT.commit, commit)
}

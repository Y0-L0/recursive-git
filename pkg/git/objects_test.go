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
	obj, err := getObject(TEST_REPO_BASE, EXAMPLE_COMMIT.sha)
	testutils.Ok(t, err)

	testutils.Equals(t, EXAMPLE_COMMIT.object, obj)
}

func TestGetCommit(t *testing.T) {
	commit, err := GetCommit(TEST_REPO_BASE, EXAMPLE_COMMIT.sha)
	testutils.Ok(t, err)

	testutils.Equals(t, &EXAMPLE_COMMIT.commit, commit)
}

var parentTests = []struct {
  id string
  sha GitSha
  parent GitSha
}{
	{
		"Normal Commit",
  	GitSha("22950c7aaaf4b990a1f69388f06a003a1462642d"),
  	GitSha("5463cfb060336eb1c6328e6ac44cf4a68779e365"),
	},
  // TODO: implement merge commit handling
  // {
  // "Merge Commit 1",
  //  GitSha("b91435bba4bba776634622252b3793afcb711910"),
  //   GitSha("22950c7aaaf4b990a1f69388f06a003a1462642d"),
  // },
  // {
  // "Merge Commit 2",
  //  GitSha("2c6bd14b0015249b232685b50ab69016e74cc775"),
  //   GitSha("153b856314764c5c4adada76156e2ef659539855"),
  // },
}
func TestGetParent(t *testing.T) {
  for _, tt := range parentTests {
    t.Run(tt.id, func(t *testing.T) {
			correctParent(t, tt.sha, tt.parent)
		})
  }
}

func correctParent(t *testing.T, sha GitSha, parent GitSha) {
	commit, err := GetCommit(TEST_REPO_BASE, sha)
	testutils.Ok(t, err)

	testutils.Equals(t, parent, commit.parent)
}

func TestGetPackedInitialCommit(t *testing.T) {
  t.Skip("Not yet implemented")
  expected := Commit{
		GitSha("b84acc25f4463b7cdaae512efdac761eac4c9c59"),
		GitSha("5463cfb060336eb1c6328e6ac44cf4a68779e365"),
    "Carlos García-Mauriño Dueñas <garcia-maurino@univention.de> 1737124337 +0100",
    "Carlos García-Mauriño Dueñas <garcia-maurino@univention.de> 1737124337 +0100",
    "Merge branch 'cgarcia/load-tests' into 'main'\n\ntest: load tests\n\nSee merge request univention/customers/dataport/upx/directory-importer!5",
  }

	commit, err := GetCommit(TEST_REPO_BASE, "6051d4147870c34253b733e6cc668055247ddb95")
	testutils.Ok(t, err)

	testutils.Equals(t, &expected, commit)
}

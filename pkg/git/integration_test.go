package git

import (
	"github.com/Y0-L0/recursive-git/testutils"
	"testing"
)

func TestGetHead(t *testing.T) {
	expected := Commit{
		tree:      "a69a5d10edd8bf796288405de4da843ce5c17238",
		parent:    "bb983db95a6067f1dbdb86d762763ad35ab8bcc2",
		author:    "@semantic-release-bot <semantic-release-bot@univention.de> 1737362940 +0000",
		committer: "@semantic-release-bot <semantic-release-bot@univention.de> 1737362940 +0000",
		message:   "chore(release): 0.1.0 [skip ci]\n\n## [0.1.0](https://git.knut.univention.de/univention/customers/dataport/upx/directory-importer/compare/v0.0.2...v0.1.0) (2025-01-20)\n\n### Features\n\n* standard nubus logging setup ([f75d623](https://git.knut.univention.de/univention/customers/dataport/upx/directory-importer/commit/f75d62306d0a8e5785b2c194817fcd4f0a3cb636))\n",
	}

	sha, err := GetHead(TEST_REPO_BASE)
	testutils.Ok(t, err)
	commit, err := GetCommit(TEST_REPO_BASE, sha)
	testutils.Ok(t, err)

	testutils.Equals(t, &expected, commit)
}

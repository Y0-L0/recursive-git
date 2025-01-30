package test

import (
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/Y0-L0/recursive-git/git/pkg/commit"
	m "github.com/Y0-L0/recursive-git/git/pkg/models"
	"github.com/Y0-L0/recursive-git/git/pkg/repo"
	"github.com/Y0-L0/recursive-git/testutils"
)

func TestMain(m *testing.M) {
	SetupLogging(slog.LevelDebug)
	code := m.Run()
	os.Exit(code)
}

func testRepo() *repo.Repo {
	return repo.NewRepo("./testdata/directory-importer/")
}

func TestGetHead(t *testing.T) {
	expected := commit.Commit{
		Tree:      "a69a5d10edd8bf796288405de4da843ce5c17238",
		Parent:    "bb983db95a6067f1dbdb86d762763ad35ab8bcc2",
		Author:    "@semantic-release-bot <semantic-release-bot@univention.de> 1737362940 +0000",
		Committer: "@semantic-release-bot <semantic-release-bot@univention.de> 1737362940 +0000",
		Message:   "chore(release): 0.1.0 [skip ci]\n\n## [0.1.0](https://git.knut.univention.de/univention/customers/dataport/upx/directory-importer/compare/v0.0.2...v0.1.0) (2025-01-20)\n\n### Features\n\n* standard nubus logging setup ([f75d623](https://git.knut.univention.de/univention/customers/dataport/upx/directory-importer/commit/f75d62306d0a8e5785b2c194817fcd4f0a3cb636))\n",
	}

	sha, err := testRepo().Head()
	testutils.Ok(t, err)
	commit, err := testRepo().Commit(sha)
	testutils.Ok(t, err)

	testutils.Equals(t, &expected, commit)
}

func TestGetCommitStack(t *testing.T) {
	expectedCommitList := []m.GitSha{
		m.GitSha("21fcd46063d09b0e178619c37bf396beece3a8e2"),
		m.GitSha("bb983db95a6067f1dbdb86d762763ad35ab8bcc2"),
		m.GitSha("b91435bba4bba776634622252b3793afcb711910"),
		// TODO: implement merge commit handling
		// m.GitSha("22950c7aaaf4b990a1f69388f06a003a1462642d"),
		// m.GitSha("5463cfb060336eb1c6328e6ac44cf4a68779e365"),
		m.GitSha("f75d62306d0a8e5785b2c194817fcd4f0a3cb636"),
		m.GitSha("9ddf53c84ad5316dc2aaf2aebedf84bbb3024169"),
		m.GitSha("6618d60463ce243f51127c3fe8ee16c960c93e07"),
		m.GitSha("180de10c4dbe4ed55efd52d5ff9d123a688f3d95"),
		m.GitSha("d9f9b1712cff9c17f82118e8e40ceb29ceeb1187"),
		m.GitSha("e3164dbbad639e801183bb01a02ee7f356134644"),
		m.GitSha("eeea1494ed65e09bb20d43bd3fc384a3e65da43a"),
		m.GitSha("2c6bd14b0015249b232685b50ab69016e74cc775"),
		// TODO: implement merge commit handling
		// m.GitSha("153b856314764c5c4adada76156e2ef659539855"),
		// m.GitSha("12d779c36d15b6c3ad10933d2feff359ed621795"),
		m.GitSha("b0278993e6530a18de832a4a672ffa901b020553"),
		m.GitSha("7e3f50f0b3489056852f87ca692aff56e31f2922"),
		m.GitSha("8aff7eb8b26d6d8ee4ff000ef5a7da139aec4638"),
		m.GitSha("f1c183375131ac0df9ab7117d54fa97c75792a25"),
		m.GitSha("edfaf701c61e1c85afdd04358685f8bfd1ef4cc9"),
		m.GitSha("e3a6e7df49d6a563a45e80e94885d564d5794ec8"),
		m.GitSha("9e807078c6dde2dfff8cd5d7f16ee2a6a3ed4944"),
		m.GitSha("715c7c0717434e6251bd9c1a66a9796a3d999c6b"),
		// TODO: Implement support for semi-linear merges
		// m.GitSha("12d9f3894581c4a31edfe49bfe30ff7f29bc212e"),
		m.GitSha("1f55ce276125be7990361281b72cb18dc69bea45"),
		m.GitSha("93d3dacf4fa8247b0218080ccc85111301886ea5"),
		m.GitSha("e2f8418556df9bbae0e8f252865140923098947a"),
		m.GitSha("abfb8c5b2df141352f2e74fc0ef6fefd134318e1"),
		m.GitSha("64c452f717e382b8cdea9a39d36506a089ce6323"),
		m.GitSha("d51b7e1c082e722c25560e9819a72d78937e92c9"),
		m.GitSha("2844c9effe0f86f0c679619f6da2616dd223df2d"),
		m.GitSha("4b235154fe74226bf68d54b3fb59a29ceb3c589c"),
		m.GitSha("9004c6fdd94939b26280264a6491ebfeb05d19f1"),
		m.GitSha("35fbd1377aa6d7f29722d62e352c75747d5ac9d1"),
		m.GitSha("a4b2059b562cd6a728d4d06ef16466ddd259b402"),
		m.GitSha("8f1733798fe95d0a9a5bd267776d1d9d0c3fc6b2"),
		m.GitSha("238af50da18daa2d463f4bdaef837bb699565f79"),
		m.GitSha("73dbba3bb3019647cdd11c58c14880644a28d25a"),
		m.GitSha("b8c75b06f333bab05e895331f0b3c50853c27c6b"),
		// m.GitSha("6051d4147870c34253b733e6cc668055247ddb95"),
	}

	sha, err := testRepo().Head()
	fmt.Println(sha)
	testutils.Ok(t, err)
	commit, err := testRepo().Commit(sha)
	testutils.Ok(t, err)

	commitList := []m.GitSha{sha}
	commitMap := map[m.GitSha]*Commit{sha: commit}

	// for commit.parent != "" {
	for commit.parent != "6051d4147870c34253b733e6cc668055247ddb95" {
		sha = commit.parent
		fmt.Println(sha)
		commit, err = testRepo().Commit(sha)
		testutils.Ok(t, err)
		commitList = append(commitList, sha)
		commitMap[sha] = commit
	}
	fmt.Println("")
	fmt.Println(commitList)
	fmt.Println(commitMap[EXAMPLE_COMMIT.sha])
	testutils.Equals(t, &EXAMPLE_COMMIT.commit, commitMap[EXAMPLE_COMMIT.sha])
	testutils.Equals(t, expectedCommitList, commitList)
}

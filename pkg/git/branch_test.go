package git

import (
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/Y0-L0/recursive-git/testutils"
)

func TestMain(m *testing.M) {
	SetupLogging(slog.LevelInfo)
	code := m.Run()
	os.Exit(code)
	slog.Debug("tests starting")
}

func TestGetHead(t *testing.T) {
	expected := Commit{
		tree:      "a69a5d10edd8bf796288405de4da843ce5c17238",
		parent:    "bb983db95a6067f1dbdb86d762763ad35ab8bcc2",
		author:    "@semantic-release-bot <semantic-release-bot@univention.de> 1737362940 +0000",
		committer: "@semantic-release-bot <semantic-release-bot@univention.de> 1737362940 +0000",
		message:   "chore(release): 0.1.0 [skip ci]\n\n## [0.1.0](https://git.knut.univention.de/univention/customers/dataport/upx/directory-importer/compare/v0.0.2...v0.1.0) (2025-01-20)\n\n### Features\n\n* standard nubus logging setup ([f75d623](https://git.knut.univention.de/univention/customers/dataport/upx/directory-importer/commit/f75d62306d0a8e5785b2c194817fcd4f0a3cb636))\n",
	}

	sha, err := testRepo().Head()
	testutils.Ok(t, err)
	commit, err := testRepo().Commit(sha)
	testutils.Ok(t, err)

	testutils.Equals(t, &expected, commit)
}

func TestResolveMergeCommits(t *testing.T) {
	expected := []GitSha{
		GitSha("bb983db95a6067f1dbdb86d762763ad35ab8bcc2"),
		GitSha("b91435bba4bba776634622252b3793afcb711910"),
		// TODO: Wrong order
		GitSha("f75d62306d0a8e5785b2c194817fcd4f0a3cb636"),
		GitSha("9ddf53c84ad5316dc2aaf2aebedf84bbb3024169"),
		GitSha("22950c7aaaf4b990a1f69388f06a003a1462642d"),
		GitSha("5463cfb060336eb1c6328e6ac44cf4a68779e365"),
		// GitSha("f75d62306d0a8e5785b2c194817fcd4f0a3cb636"),
		// GitSha("9ddf53c84ad5316dc2aaf2aebedf84bbb3024169"),
	}

	branch := newBranch(testRepo(), expected[0])

	err := branch.Resolve()
	testutils.Ok(t, err)

	fmt.Println(len(branch.List))
	testutils.Equals(t, 39, len(branch.List))
	testutils.Equals(t, expected, branch.List[:6])
}

var expectedCommitList = []GitSha{
	GitSha("21fcd46063d09b0e178619c37bf396beece3a8e2"),
	GitSha("bb983db95a6067f1dbdb86d762763ad35ab8bcc2"),
	GitSha("b91435bba4bba776634622252b3793afcb711910"),
	// TODO: Wrong order
	// GitSha("22950c7aaaf4b990a1f69388f06a003a1462642d"),
	// GitSha("5463cfb060336eb1c6328e6ac44cf4a68779e365"),
	GitSha("f75d62306d0a8e5785b2c194817fcd4f0a3cb636"),
	GitSha("9ddf53c84ad5316dc2aaf2aebedf84bbb3024169"),
	GitSha("22950c7aaaf4b990a1f69388f06a003a1462642d"),
	GitSha("5463cfb060336eb1c6328e6ac44cf4a68779e365"),
	GitSha("6618d60463ce243f51127c3fe8ee16c960c93e07"),
	GitSha("180de10c4dbe4ed55efd52d5ff9d123a688f3d95"),
	GitSha("d9f9b1712cff9c17f82118e8e40ceb29ceeb1187"),
	GitSha("e3164dbbad639e801183bb01a02ee7f356134644"),
	GitSha("eeea1494ed65e09bb20d43bd3fc384a3e65da43a"),
	GitSha("2c6bd14b0015249b232685b50ab69016e74cc775"),
	// TODO: implement merge commit handling
	GitSha("153b856314764c5c4adada76156e2ef659539855"),
	GitSha("12d779c36d15b6c3ad10933d2feff359ed621795"),
	GitSha("b0278993e6530a18de832a4a672ffa901b020553"),
	GitSha("7e3f50f0b3489056852f87ca692aff56e31f2922"),
	GitSha("8aff7eb8b26d6d8ee4ff000ef5a7da139aec4638"),
	GitSha("f1c183375131ac0df9ab7117d54fa97c75792a25"),
	GitSha("edfaf701c61e1c85afdd04358685f8bfd1ef4cc9"),
	GitSha("e3a6e7df49d6a563a45e80e94885d564d5794ec8"),
	GitSha("9e807078c6dde2dfff8cd5d7f16ee2a6a3ed4944"),
	GitSha("715c7c0717434e6251bd9c1a66a9796a3d999c6b"),
	// TODO: Implement support for semi-linear merges
	GitSha("12d9f3894581c4a31edfe49bfe30ff7f29bc212e"),
	GitSha("1f55ce276125be7990361281b72cb18dc69bea45"),
	GitSha("93d3dacf4fa8247b0218080ccc85111301886ea5"),
	GitSha("e2f8418556df9bbae0e8f252865140923098947a"),
	GitSha("abfb8c5b2df141352f2e74fc0ef6fefd134318e1"),
	GitSha("64c452f717e382b8cdea9a39d36506a089ce6323"),
	GitSha("d51b7e1c082e722c25560e9819a72d78937e92c9"),
	GitSha("2844c9effe0f86f0c679619f6da2616dd223df2d"),
	GitSha("4b235154fe74226bf68d54b3fb59a29ceb3c589c"),
	GitSha("9004c6fdd94939b26280264a6491ebfeb05d19f1"),
	GitSha("35fbd1377aa6d7f29722d62e352c75747d5ac9d1"),
	GitSha("a4b2059b562cd6a728d4d06ef16466ddd259b402"),
	GitSha("8f1733798fe95d0a9a5bd267776d1d9d0c3fc6b2"),
	GitSha("238af50da18daa2d463f4bdaef837bb699565f79"),
	GitSha("73dbba3bb3019647cdd11c58c14880644a28d25a"),
	GitSha("b8c75b06f333bab05e895331f0b3c50853c27c6b"),
	GitSha("6051d4147870c34253b733e6cc668055247ddb95"),
}

func TestResolveHeadBranch(t *testing.T) {
	branch, err := testRepo().HeadBranch()
	testutils.Ok(t, err)
	err = branch.Resolve()
	testutils.Ok(t, err)

	testutils.Equals(t, true, branch.In(EXAMPLE_COMMIT.sha))
	testutils.Equals(t, expectedCommitList, branch.List)
}

func TestResolveBranch(t *testing.T) {
	branch, err := testRepo().Branch("main")
	testutils.Ok(t, err)
	err = branch.Resolve()
	testutils.Ok(t, err)

	testutils.Equals(t, true, branch.In(EXAMPLE_COMMIT.sha))
	testutils.Equals(t, expectedCommitList, branch.List)
}

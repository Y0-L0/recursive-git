package git

func (suite *GitTest) TestParseCommit() {
	commit, err := newCommit(EXAMPLE_COMMIT.object)
	suite.NoError(err)

	suite.Equal(&EXAMPLE_COMMIT.commit, commit)
}

func (suite *GitTest) TestGetCommit() {
	commit, err := testRepo().Commit(EXAMPLE_COMMIT.sha)
	suite.NoError(err)

	suite.Equal(&EXAMPLE_COMMIT.commit, commit)
}

func (suite *GitTest) TestGetMergeCommit() {
	expectedParent := GitSha("22950c7aaaf4b990a1f69388f06a003a1462642d")
	expectedMergeParent := GitSha("b91435bba4bba776634622252b3793afcb711910")
	commit, err := testRepo().Commit(GitSha("bb983db95a6067f1dbdb86d762763ad35ab8bcc2"))
	suite.NoError(err)

	suite.Equal(expectedParent, commit.parent)
	suite.Equal(expectedMergeParent, commit.mergeParent)
}

// var parentTests = []struct {
// 	id     string
// 	sha    GitSha
// 	parent GitSha
// }{
// 	{
// 		"Normal Commit",
// 	},
// 	// TODO: implement merge commit handling
// 	// {
// 	// "Merge Commit 1",
// 	//  GitSha("b91435bba4bba776634622252b3793afcb711910"),
// 	//  GitSha("22950c7aaaf4b990a1f69388f06a003a1462642d"),
// 	// },
// 	// {
// 	// "Merge Commit 2",
// 	//  GitSha("2c6bd14b0015249b232685b50ab69016e74cc775"),
// 	//  GitSha("153b856314764c5c4adada76156e2ef659539855"),
// 	// },
// }
//
// func (suite *GitTest) TestGetParent() {
// 	for _, tt := range parentTests {
// 		t.Run(tt.id, func(t *testing.T) {
// 			correctParent(t, tt.sha, tt.parent)
// 		})
// 	}
// }

func (suite *GitTest) TestGetParent() {
	sha := GitSha("22950c7aaaf4b990a1f69388f06a003a1462642d")
	parent := GitSha("6618d60463ce243f51127c3fe8ee16c960c93e07")
	commit, err := testRepo().Commit(sha)
	suite.NoError(err)

	suite.Equal(parent, commit.parent)
}

func (suite *GitTest) TestGetPackedInitialCommit() {
	suite.T().Skip("Not yet implemented")
	expected := Commit{
		GitSha("b84acc25f4463b7cdaae512efdac761eac4c9c59"),
		GitSha("5463cfb060336eb1c6328e6ac44cf4a68779e365"),
		GitSha(""),
		"Carlos García-Mauriño Dueñas <garcia-maurino@univention.de> 1737124337 +0100",
		"Carlos García-Mauriño Dueñas <garcia-maurino@univention.de> 1737124337 +0100",
		0,
		"Merge branch 'cgarcia/load-tests' into 'main'\n\ntest: load tests\n\nSee merge request univention/customers/dataport/upx/directory-importer!5",
	}

	commit, err := testRepo().Commit("6051d4147870c34253b733e6cc668055247ddb95")
	suite.NoError(err)

	suite.Equal(&expected, commit)
}

package git

func testRepo() *Repo {
	return NewRepo("./testdata/directory-importer/")
}

func (suite *GitTest) TestGetHeadSuccess() {
	sha, err := testRepo().Head()
	suite.NoError(err)

	suite.Equal(GitSha("21fcd46063d09b0e178619c37bf396beece3a8e2"), sha)
}

func (suite *GitTest) TestGetHeadNoFile() {
	sha, err := testRepo().ref("refs/invalid-location-heads/main")
	suite.Equal(GitSha(""), sha)

	// TODO: Validate that the correct error is returned!
	suite.Error(err, "Expeced a file not found error but got nil")
}

func (suite *GitTest) TestGetRefSuccess() {
	sha, err := testRepo().ref("refs/heads/main")
	suite.NoError(err)

	suite.Equal(GitSha("21fcd46063d09b0e178619c37bf396beece3a8e2"), sha)
}

func (suite *GitTest) TestGetRefNoFile() {
	sha, err := testRepo().ref("refs/invalid-location-heads/main")
	suite.Equal(GitSha(""), sha)

	// TODO: Validate that the correct error is returned!
	suite.Error(err, "Expeced a file not found error but got nil")
}

func (suite *GitTest) TestGetRefWrongContent() {
	sha, err := testRepo().ref("HEAD")
	suite.Equal(GitSha(""), sha)

	// TODO: Validate that the correct error is returned!
	suite.Error(err, "Expected an InvalidContent error but got nil")
}

func (suite *GitTest) TestCommitCaching() {
	repo := testRepo()
	commit, err := repo.Commit(EXAMPLE_COMMIT.sha)
	suite.NoError(err)

	commit1, err := repo.Commit(EXAMPLE_COMMIT.sha)
	suite.NoError(err)

	if commit != commit1 {
		suite.Fail("\nexp: %p\ngot: %p", commit, commit1)
	}
}

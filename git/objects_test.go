package git

var EXAMPLE_COMMIT = struct {
	sha    GitSha
	object string
	commit Commit
}{
	sha:    GitSha("eeea1494ed65e09bb20d43bd3fc384a3e65da43a"),
	object: "commit 324\x00tree c9be4a262387887608b3c3c3d1237b43f5d3ac82\nparent 2c6bd14b0015249b232685b50ab69016e74cc775\nauthor Johannes Lohmer <lohmer@univention.de> 1736442888 +0100\ncommitter Johannes Lohmer <lohmer@univention.de> 1736443347 +0100\n\ntest: Create and delete the example.org maildomain as part of every testrun with an autouse fixture\n",
	commit: Commit{
		GitSha("eeea1494ed65e09bb20d43bd3fc384a3e65da43a"),
		GitSha("c9be4a262387887608b3c3c3d1237b43f5d3ac82"),
		GitSha("2c6bd14b0015249b232685b50ab69016e74cc775"),
		GitSha(""),
		"Johannes Lohmer <lohmer@univention.de> 1736442888 +0100",
		"Johannes Lohmer <lohmer@univention.de> 1736443347 +0100",
		1736443347,
		"test: Create and delete the example.org maildomain as part of every testrun with an autouse fixture\n",
	},
}

func (suite *GitTest) TestGetObject() {
	obj, err := getObject(testRepo().base, EXAMPLE_COMMIT.sha)
	suite.NoError(err)

	suite.Equal(EXAMPLE_COMMIT.object, obj)
}

package git

type GitSha string

type Commit struct {
	tree      GitSha
	parent    GitSha
	author    string
	committer string
	message   string
}

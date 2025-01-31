package git

type Branch struct {
	head GitSha
	repo *Repo
	list []GitSha
	set  map[GitSha]bool
}

func newBranch(repo *Repo, head GitSha) *Branch {
	return &Branch{
		head: head,
		repo: repo,
		set:  make(map[GitSha]bool),
	}
}

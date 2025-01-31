package git

import "log/slog"

type Branch struct {
	head GitSha
	repo *Repo
	set  map[GitSha]bool
	List []GitSha
}

func newBranch(repo *Repo, head GitSha) *Branch {
	return &Branch{
		head: head,
		repo: repo,
		set:  map[GitSha]bool{head: true},
		List: []GitSha{head},
	}
}

func (branch *Branch) In(sha GitSha) bool {
	return branch.set[sha]
}

func (branch *Branch) Resolve() error {
	commit, err := branch.repo.Commit(branch.head)
	if err != nil {
		return err
	}

	// TODO: Implement packed commit reading
	// for commit.parent != "" {
	for commit.parent != "6051d4147870c34253b733e6cc668055247ddb95" {
		sha := commit.parent
		commit, err = branch.repo.Commit(sha)
		if err != nil {
			return err
		}
		branch.List = append(branch.List, sha)
		branch.set[sha] = true
	}
	slog.Debug("successfully resoved git branch", "commitList", branch.List)
	return nil
}

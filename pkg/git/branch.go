package git

import (
	"log/slog"
)

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
		set:  map[GitSha]bool{},
		List: []GitSha{head},
	}
}

func (branch *Branch) In(sha GitSha) bool {
	return branch.set[sha]
}

func (branch *Branch) Resolve() error {
	list, err := branch.getParents(branch.head)
	if err != nil {
		return err
	}
	slog.Debug("successfully resoved git branch", "commitList", list)
	branch.List = list
	return nil
}

func (branch *Branch) getParents(sha GitSha) ([]GitSha, error) {
	slog.Debug("resolved parent", "sha", sha)

	if branch.set[sha] {
		slog.Info("Found sha already in the branch.", "sha", sha)
		return []GitSha{}, nil
	}
	branch.set[sha] = true
	if sha == "6051d4147870c34253b733e6cc668055247ddb95" {
		return []GitSha{sha}, nil
	}

	commit, err := branch.repo.Commit(sha)
	if err != nil {
		return nil, err
	}

	parents, err := branch.getParents(commit.parent)
	if err != nil {
		return nil, err
	}
	var mergeParents []GitSha

	if commit.mergeParent != GitSha("") {
		mergeParents, err = branch.getParents(commit.mergeParent)
		if err != nil {
			return nil, err
		}
	}
	slog.Warn("resolved parents", "sha", sha, "mergeParents", mergeParents, "parents", parents)

	result := []GitSha{sha}
	result = append(result, mergeParents...)
	result = append(result, parents...)

	return result, nil
}

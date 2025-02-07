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
	result = append(result, branch.mergeParentArrays(parents, mergeParents)...)

	return result, nil
}

func (branch *Branch) mergeParentArrays(parents []GitSha, mergeParents []GitSha) []GitSha {
	if len(mergeParents) == 0 {
		return parents
	}

	parentIndex := 0
	maxParentIndex := len(parents)
	mergeParentIndex := 0
	maxMergeParentIndex := len(mergeParents)

	var result []GitSha

	for parentIndex < maxParentIndex && mergeParentIndex < maxMergeParentIndex {
		parentCommit, err := branch.repo.Commit(parents[parentIndex])
		if err != nil {
			panic("Git commit database or cache got corrupted")
		}
		mergeParentCommit, err := branch.repo.Commit(parents[parentIndex])
		if err != nil {
			panic("Git commit database or cache got corrupted")
		}

		if parentCommit.committerEpoch < mergeParentCommit.committerEpoch {
			result = append(result, parents[parentIndex])
			parentIndex++
		} else {
			result = append(result, mergeParents[mergeParentIndex])
			mergeParentIndex++
		}
	}
	if parentIndex == maxParentIndex {
		result = append(result, mergeParents[mergeParentIndex:]...)
	} else if mergeParentIndex == maxMergeParentIndex {
		result = append(result, parents[parentIndex:]...)
	} else {
		slog.Error("Logic error in mergeParentArrays", "parentIndex", parentIndex, "maxParentIndex", maxParentIndex, "mergeParentIndex", mergeParentIndex, "maxMergeParentIndex", maxMergeParentIndex, "parents", parents, "mergeParents", mergeParents)
		panic("impossible ")
	}

	slog.Debug("Final state of mergeParents", "parentIndex", parentIndex, "maxParentIndex", maxParentIndex, "mergeParentIndex", mergeParentIndex, "maxMergeParentIndex", maxMergeParentIndex, "parents", parents, "mergeParents", mergeParents)
	slog.Info("Merged parents", "result", result)
	return result
}

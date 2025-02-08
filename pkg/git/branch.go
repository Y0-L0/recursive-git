package git

import (
	"iter"
	"log/slog"
	"slices"
	"sort"
)

type commitSlice []*Commit

func (cs commitSlice) Insert(commit *Commit) commitSlice {
	index := sort.Search(len(cs), func(i int) bool {
		return cs[i].committerEpoch > commit.committerEpoch
	})
	return slices.Insert(cs, index, commit)
}

func (cs commitSlice) Pop() (commitSlice, *Commit) {
	result := cs[len(cs)-1]
	return cs[:len(cs)-1], result
}

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

func (branch *Branch) ResolveIterator() error {
	list, err := branch.getParentsRecursively(branch.head)
	if err != nil {
		return err
	}
	slog.Debug("successfully resoved git branch using iterators", "commitList", list)
	branch.List = list
	return nil
}

func (branch *Branch) ResolveRecursively() error {
	list, err := branch.getParentsRecursively(branch.head)
	if err != nil {
		return err
	}
	slog.Debug("successfully resoved git branch using recursion", "commitList", list)
	branch.List = list
	return nil
}

func (branch *Branch) parents(commit *Commit) iter.Seq2[GitSha, error] {
	stack := commitSlice{commit}
	return func(yield func(GitSha, error) bool) {
		for len(stack) != 0 {
			stack, currentCommit := stack.Pop()
			slog.Debug("Removed element from slice", "new length", len(stack))
			parent, err := branch.repo.Commit(currentCommit.parent)
			if err != nil {
				if !yield(GitSha(""), err) {
					return
				}
			}
			stack = stack.Insert(parent)
			slog.Debug("Inserted element into slice", "new length", len(stack))
			if currentCommit.mergeParent != GitSha("") {
				mergeParent, err := branch.repo.Commit(commit.mergeParent)
				if err != nil {
					if !yield(GitSha(""), err) {
						return
					}
				}
				stack = stack.Insert(mergeParent)
				slog.Debug("Inserted element into slice", "new length", len(stack))
			}
			if !yield(stack[len(stack)-1].sha, nil) {
				return
			}
		}
	}
}

func (branch *Branch) getParentsRecursively(sha GitSha) ([]GitSha, error) {
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

	parents, err := branch.getParentsRecursively(commit.parent)
	if err != nil {
		return nil, err
	}
	var mergeParents []GitSha

	if commit.mergeParent != GitSha("") {
		mergeParents, err = branch.getParentsRecursively(commit.mergeParent)
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

package git

import (
	"fmt"
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
	list []GitSha
}

func newBranch(repo *Repo, head GitSha) *Branch {
	return &Branch{
		head: head,
		repo: repo,
		set:  map[GitSha]bool{},
		list: []GitSha{},
	}
}

func (branch *Branch) In(sha GitSha) bool {
	return branch.set[sha]
}

func (branch *Branch) List() ([]GitSha, error) {
	if len(branch.list) != 0 {
		return branch.list, nil
	}

	commit, err := branch.repo.Commit(branch.head)
	if err != nil {
		return nil, err
	}

	list := []GitSha{branch.head}
	for sha, err := range branch.Iterate(commit) {
		if err != nil {
			return nil, err
		}
		list = append(list, sha)
	}

	slog.Debug("successfully resoved git branch using iterators", "commitList", list)
	branch.list = list
	return list, nil
}

func (branch *Branch) Iterate(commit *Commit) iter.Seq2[GitSha, error] {
	stack := commitSlice{commit}
	counter := 0

	resolveParent := func(sha GitSha) error {
		if sha != "" && !branch.set[sha] {
			branch.set[sha] = true
			commit, err := branch.repo.Commit(sha)
			if err != nil {
				return err
			}
			stack = stack.Insert(commit)
			slog.Debug("Inserted (merge) parent into commit stack", "new length", len(stack), "sha", sha, "stack", stack)
		}
		return nil
	}

	return func(yield func(GitSha, error) bool) {
		for len(stack) != 0 {
			counter++
			if counter >= 50 {
				slog.Error("infinite loop protection")
				return
			}

			currentCommit := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			slog.Debug("Removed element from commit stack", "new length", len(stack), "currentCommit", currentCommit.sha, "stack", stack)

			err := resolveParent(currentCommit.parent)
			if err != nil {
				if !yield(GitSha(""), err) {
					return
				}
			}

			err = resolveParent(currentCommit.mergeParent)
			if err != nil {
				if !yield(GitSha(""), err) {
					return
				}
			}

			if len(stack) == 0 {
				fmt.Print("Finished")
				return
			}
			if !yield(stack[len(stack)-1].sha, nil) {
				return
			}
		}
	}
}

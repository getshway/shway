package git

import (
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// Clone a repository into the path with the given options
func Clone(path, u string) (commit *object.Commit, err error) {
	r, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:               u,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	// Retrieving the commit object
	ref, err := r.Head()
	if err != nil {
		return
	}
	return r.CommitObject(ref.Hash())
}

// Pull changes from a remote repository
func Pull(path string) (commit *object.Commit, err error) {
	// Instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(path)
	if err != nil {
		return
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		return
	}

	// Pull the latest changes from the origin remote and merge into the current branch
	if err = w.Pull(&git.PullOptions{RemoteName: "origin"}); err != nil {
		return
	}

	// Get latest commit that was just pulled
	ref, err := r.Head()
	if err != nil {
		return
	}
	return r.CommitObject(ref.Hash())
}

# git-merged-master

[![Build Status](https://travis-ci.org/petems/git-merged-master.svg?branch=master)](https://travis-ci.org/petems/git-merged-master)

A helper function that uses the `src-d/go-git.v4` library is a determine what branches have already been merged into master in a remote

It iterates through the commits in master, and the head branch of all existing remote branches, and if the head branch of the remote branches matches the commit in master then it is marked as a merged branch:

```golang
err = masterCommits.ForEach(func(commit *object.Commit) error {
	for branchName, branchHead := range remoteBranchHeads {
	  if branchHead.String() == commit.Hash.String() {
			log.Infof("Branch %s head (%s) was found in master, so has been merged!\n", branchName, branchHead)
			mergedBranches = append(mergedBranches, branchName)
		}
	}
	return nil
})
```

## Usage

WIP
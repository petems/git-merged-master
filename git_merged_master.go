package gitmergedmaster

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"

	log "github.com/sirupsen/logrus"
)

// IsStringInSlice checks if a string is present in a slice of strings
//
//	letters := []string{"a", "b", "c", "d"}
//  IsStringInSlice("a", letters) // true
//  IsStringInSlice("e", letters) // false
//
func isStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func remoteBranchesToStrings(gitRemoteArray []*git.Remote) []string {
	stringArray := make([]string, len(gitRemoteArray))
	for i, v := range gitRemoteArray {
		stringArray[i] = v.Config().Name
	}
	return stringArray
}

// getRemoteBranches returns a storer.ReferenceIter of all the branches that are set to a remote
func getRemoteBranches(s storer.ReferenceStorer) (storer.ReferenceIter, error) {
	refs, err := s.IterReferences()
	if err != nil {
		return nil, err
	}

	return storer.NewReferenceFilteredIter(func(ref *plumbing.Reference) bool {
		return ref.Name().IsRemote()
	}, refs), nil
}

func getMergedBranches(gitRepo *git.Repository, remoteOrigin, masterBranchName, skipBranches string) ([]string, error) {

	log.Info("Attempting to get master information from branches from repo")

	branchRefs, err := gitRepo.Branches()
	if err != nil {
		log.Fatal("list branches failed:", err)
		return nil, err
	}

	branchHeads := make(map[string]plumbing.Hash)

	log.Info("Fetching from the remote...")

	listRemotes, err := gitRepo.Remotes()

	if err != nil {
		log.Fatal("Error looking for remotes", err)
		return nil, err
	}

	remoteBranchesAsStrings := remoteBranchesToStrings(listRemotes)

	if !isStringInSlice(remoteOrigin, remoteBranchesAsStrings) {
		return nil, errors.New("Could not find the remote named " + remoteOrigin)
	}

	err = branchRefs.ForEach(func(reference *plumbing.Reference) error {
		branchName := strings.TrimPrefix(reference.Name().String(), "refs/heads/")
		branchHead := reference.Hash()
		branchHeads[branchName] = branchHead
		return nil
	})

	if err != nil {
		log.Fatal("list branches failed:", err)
		return nil, err
	}

	masterCommits, err := gitRepo.Log(&git.LogOptions{From: branchHeads[masterBranchName]})

	if err != nil {
		log.Fatal("get commits from master failed:", err)
		return nil, err
	}

	remoteBranches, err := getRemoteBranches(gitRepo.Storer)

	if err != nil {
		log.Fatal("list remote branches failed:", err)
		return nil, err
	}

	remoteBranchHeads := make(map[string]plumbing.Hash)

	err = remoteBranches.ForEach(func(branch *plumbing.Reference) error {
		remoteBranchName := strings.TrimPrefix(branch.Name().String(), "refs/remotes/")
		remoteBranchHead := branch.Hash()
		remoteBranchHeads[remoteBranchName] = remoteBranchHead
		return nil
	})

	if err != nil {
		log.Fatal("iterating remote branches failed:", err)
		return nil, err
	}

	for branchName, branchHead := range remoteBranchHeads {
		log.Infof("Remote Branch %s head is: %s", branchName, branchHead)
	}

	mergedBranches := make([]string, 0)

	masterBranchRemote := fmt.Sprintf("%s/%s", remoteOrigin, masterBranchName)

	delete(remoteBranchHeads, masterBranchRemote)

	err = masterCommits.ForEach(func(commit *object.Commit) error {
		for branchName, branchHead := range remoteBranchHeads {
			if branchHead.String() == commit.Hash.String() {
				log.Infof("Branch %s head (%s) was found in master, so has been merged!\n", branchName, branchHead)
				mergedBranches = append(mergedBranches, branchName)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal("looking for merged commits failed:", err)
		return nil, err
	}

	sort.Strings(mergedBranches)

	return mergedBranches, nil
}

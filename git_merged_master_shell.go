package gitmergedmaster

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	gitshell "code.gitea.io/git"
)

func getMergedBranchesShell(gitRepo, remoteOrigin, masterBranchName, skipBranches string) ([]string, error) {

	_, err := gitshell.NewCommand("ls-remote", "--exit-code", "--heads", remoteOrigin).RunInDir(gitRepo)

	if err != nil {
		return nil, errors.New("Could not find the remote named " + remoteOrigin)
	}

	_, err = gitshell.NewCommand("ls-remote", "--exit-code", "--heads", remoteOrigin, masterBranchName).RunInDir(gitRepo)

	if err != nil {
		return nil, errors.New("Could not find the master branch " + masterBranchName + " on remote " + remoteOrigin)
	}

	remoteMaster := fmt.Sprintf("%s/%s", remoteOrigin, masterBranchName)
	remoteHead := fmt.Sprintf("%s/%s", remoteOrigin, "HEAD")

	shellMergedBranchesOutput, err := gitshell.NewCommand("branch", "-r", "--format", "'%(refname:short)'", "--merged", remoteMaster).RunInDir(gitRepo)

	shellMergedBranchesOutputQuoteless := strings.Replace(shellMergedBranchesOutput, "'", "", -1)
	shellMergedBranchesOutputMasterless := strings.Replace(shellMergedBranchesOutputQuoteless, remoteMaster, "", 1)
	shellMergedBranchesOutputHeadless := strings.Replace(shellMergedBranchesOutputMasterless, remoteHead, "", 1)

	branchesSlice := strings.Fields(shellMergedBranchesOutputHeadless)

	sort.Strings(branchesSlice)

	return branchesSlice, err
}

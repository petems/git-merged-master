package gitmergedmaster

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	gitshell "code.gitea.io/git"
)

func setupRepoWithLocalRemote() string {
	repoName := fmt.Sprintf("repo-tmp-golang-%s", time.Now())

	tmpDir, err := ioutil.TempDir("", repoName)

	if err != nil {
		panic(err)
	}

	err = gitshell.InitRepository(tmpDir, false)

	if err != nil {
		panic(err)
	}

	_, err = os.Create(filepath.Join(tmpDir, "test.txt"))

	if err != nil {
		panic(err)
	}

	err = gitshell.AddChanges(tmpDir, true)

	if err != nil {
		panic(err)
	}

	err = gitshell.CommitChanges(tmpDir, gitshell.CommitChangesOptions{
		Committer: &gitshell.Signature{
			Email: "user2@example.com",
			Name:  "User Two",
			When:  time.Now(),
		},
		Author: &gitshell.Signature{
			Email: "user2@example.com",
			Name:  "User Two",
			When:  time.Now(),
		},
		Message: fmt.Sprintf("Testing push create @ %v", time.Now()),
	})

	if err != nil {
		panic(err)
	}

	_, err = gitshell.NewCommand("remote", "add", "local", ".").RunInDir(tmpDir)

	gitshell.NewCommand("push", "-u", "local", "master").RunInDir(tmpDir)

	if err != nil {
		panic(err)
	}

	return tmpDir
}

func setupRepoWithLocalRemoteWithNCommits(commits int) string {
	repoName := fmt.Sprintf("repo-tmp-golang-%s", time.Now())

	tmpDir, err := ioutil.TempDir("", repoName)

	if err != nil {
		panic(err)
	}

	err = gitshell.InitRepository(tmpDir, false)

	if err != nil {
		panic(err)
	}

	_, err = os.Create(filepath.Join(tmpDir, "test.txt"))

	if err != nil {
		panic(err)
	}

	err = gitshell.AddChanges(tmpDir, true)

	if err != nil {
		panic(err)
	}

	for i := 0; i < commits; i++ {
		err = gitshell.CommitChanges(tmpDir, gitshell.CommitChangesOptions{
			Committer: &gitshell.Signature{
				Email: "user2@example.com",
				Name:  "User Two",
				When:  time.Now(),
			},
			Author: &gitshell.Signature{
				Email: "user2@example.com",
				Name:  "User Two",
				When:  time.Now(),
			},
			Message: fmt.Sprintf("Testing push create @ %v", time.Now()),
		})

		if err != nil {
			panic(err)
		}
	}

	_, err = gitshell.NewCommand("remote", "add", "local", ".").RunInDir(tmpDir)

	gitshell.NewCommand("push", "-u", "local", "master").RunInDir(tmpDir)

	if err != nil {
		panic(err)
	}

	return tmpDir
}

func createBranchesCount(repolocation string, branchcount int) {
	for i := 0; i < branchcount; i++ {
		branchName := "from-master-" + strconv.Itoa(i)
		createBranchName(repolocation, branchName)
	}
}

func createBranchName(repolocation string, branchname string) {
	gitshell.NewCommand("checkout", "-b", branchname).RunInDir(repolocation)
	gitshell.NewCommand("push", "-u", "local", branchname).RunInDir(repolocation)
}

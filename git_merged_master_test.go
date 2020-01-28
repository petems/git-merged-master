package gitmergedmaster

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	gitshell "code.gitea.io/git"

	"gopkg.in/src-d/go-git.v4"

	"github.com/stretchr/testify/assert"
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

func TestRepoValidRemote(t *testing.T) {
	tmpdir := setupRepoWithLocalRemote()

	repo, err := git.PlainOpen(tmpdir)

	assert.NoError(t, err)

	_, err = getMergedBranches(repo, "local", "master", "")

	assert.NoError(t, err)
}
func TestRepoInvalidRemote(t *testing.T) {
	tmpdir := setupRepoWithLocalRemote()

	repo, err := git.PlainOpen(tmpdir)

	assert.NoError(t, err)

	_, err = getMergedBranches(repo, "not-found", "master", "")

	assert.EqualError(t, err, "Could not find the remote named not-found")
}

func TestRepoNamedBranch(t *testing.T) {
	tmpdir := setupRepoWithLocalRemote()

	createBranchName(tmpdir, "named-branch")

	repo, err := git.PlainOpen(tmpdir)

	assert.NoError(t, err)

	branchArray, err := getMergedBranches(repo, "local", "master", "")

	assert.NoError(t, err)
	assert.Equal(t, branchArray[0], "local/named-branch")
	assert.Equal(t, len(branchArray), 1)
}

func TestRepo5Branches(t *testing.T) {
	tmpdir := setupRepoWithLocalRemote()

	createBranchesCount(tmpdir, 5)

	repo, err := git.PlainOpen(tmpdir)

	assert.NoError(t, err)

	branchArray, err := getMergedBranches(repo, "local", "master", "")

	assert.NoError(t, err)
	assert.Equal(t, len(branchArray), 5)
}

// BENCHMARKS

func benchmarkGetMergedBranches(branchCount int, t *testing.B) {
	tmpdir := setupRepoWithLocalRemote()

	createBranchesCount(tmpdir, branchCount)

	repo, err := git.PlainOpen(tmpdir)

	assert.NoError(t, err)

	t.ResetTimer()

	branchArray, err := getMergedBranches(repo, "local", "master", "")

	assert.NoError(t, err)
	assert.Equal(t, len(branchArray), branchCount)
}

func BenchmarkGMG1(b *testing.B)   { benchmarkGetMergedBranches(1, b) }
func BenchmarkGMG2(b *testing.B)   { benchmarkGetMergedBranches(2, b) }
func BenchmarkGMG3(b *testing.B)   { benchmarkGetMergedBranches(3, b) }
func BenchmarkGMG10(b *testing.B)  { benchmarkGetMergedBranches(10, b) }
func BenchmarkGMG20(b *testing.B)  { benchmarkGetMergedBranches(20, b) }
func BenchmarkGMG40(b *testing.B)  { benchmarkGetMergedBranches(40, b) }
func BenchmarkGMG100(b *testing.B) { benchmarkGetMergedBranches(100, b) }

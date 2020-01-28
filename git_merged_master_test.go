package gitmergedmaster

import (
	"testing"

	"gopkg.in/src-d/go-git.v4"

	"github.com/stretchr/testify/assert"
)

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

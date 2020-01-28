package gitmergedmaster

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellRepoValidRemote(t *testing.T) {
	tmpdir := setupRepoWithLocalRemote()

	_, err := getMergedBranchesShell(tmpdir, "local", "master", "")

	assert.NoError(t, err)
}
func TestShellRepoInvalidRemote(t *testing.T) {
	tmpdir := setupRepoWithLocalRemote()

	_, err := getMergedBranchesShell(tmpdir, "not-found", "master", "")

	assert.EqualError(t, err, "Could not find the remote named not-found")
}

func TestShellRepoInvalidMasterOnRemote(t *testing.T) {
	tmpdir := setupRepoWithLocalRemote()

	_, err := getMergedBranchesShell(tmpdir, "local", "master-not-exist", "")

	assert.EqualError(t, err, "Could not find the master branch master-not-exist on remote local")
}

func TestShellRepoNamedBranch(t *testing.T) {
	tmpdir := setupRepoWithLocalRemote()

	createBranchName(tmpdir, "named-branch")

	branchArray, err := getMergedBranchesShell(tmpdir, "local", "master", "")

	assert.NoError(t, err)
	assert.Equal(t, "local/named-branch", branchArray[0])
	assert.Equal(t, 1, len(branchArray))
}

func TestShellRepo5Branches(t *testing.T) {
	tmpdir := setupRepoWithLocalRemote()

	createBranchesCount(tmpdir, 5)

	branchArray, err := getMergedBranchesShell(tmpdir, "local", "master", "")

	assert.NoError(t, err)

	assert.Equal(t, "local/from-master-0", branchArray[0])
	assert.Equal(t, "local/from-master-1", branchArray[1])
	assert.Equal(t, "local/from-master-2", branchArray[2])
	assert.Equal(t, "local/from-master-3", branchArray[3])
	assert.Equal(t, "local/from-master-4", branchArray[4])

	assert.Equal(t, 5, len(branchArray))
}

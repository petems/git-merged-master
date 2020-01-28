package gitmergedmaster

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
)

func benchmarkGetMergedBranches(branchCount int, t *testing.B) {
	tmpdir := setupRepoWithLocalRemote()

	createBranchesCount(tmpdir, branchCount)

	repo, err := git.PlainOpen(tmpdir)

	assert.NoError(t, err)

	t.ResetTimer()

	for n := 0; n < t.N; n++ {
		getMergedBranches(repo, "local", "master", "")
	}

}

func benchmarkGetMergedBranchesShell(branchCount int, t *testing.B) {
	tmpdir := setupRepoWithLocalRemote()

	createBranchesCount(tmpdir, branchCount)

	t.ResetTimer()

	for n := 0; n < t.N; n++ {
		getMergedBranchesShell(tmpdir, "local", "master", "")

	}

}

func benchmarkGetMergedBranchesCommits(commits int, branchCount int, t *testing.B) {
	tmpdir := setupRepoWithLocalRemoteWithNCommits(commits)

	createBranchesCount(tmpdir, branchCount)

	repo, err := git.PlainOpen(tmpdir)

	assert.NoError(t, err)

	t.ResetTimer()

	for n := 0; n < t.N; n++ {
		getMergedBranches(repo, "local", "master", "")
	}
}

func benchmarkGetMergedBranchesCommitsShell(commits int, branchCount int, t *testing.B) {
	tmpdir := setupRepoWithLocalRemoteWithNCommits(commits)

	createBranchesCount(tmpdir, branchCount)

	t.ResetTimer()

	for n := 0; n < t.N; n++ {
		getMergedBranchesShell(tmpdir, "local", "master", "")
	}

}

func BenchmarkGMGBranches1(b *testing.B)  { benchmarkGetMergedBranches(1, b) }
func BenchmarkGMGBranches2(b *testing.B)  { benchmarkGetMergedBranches(2, b) }
func BenchmarkGMGBranches3(b *testing.B)  { benchmarkGetMergedBranches(3, b) }
func BenchmarkGMGBranches10(b *testing.B) { benchmarkGetMergedBranches(10, b) }
func BenchmarkGMGBranches20(b *testing.B) { benchmarkGetMergedBranches(20, b) }
func BenchmarkGMGBranches40(b *testing.B) { benchmarkGetMergedBranches(40, b) }

func BenchmarkGMGShellBranches1(b *testing.B)  { benchmarkGetMergedBranchesShell(1, b) }
func BenchmarkGMGShellBranches2(b *testing.B)  { benchmarkGetMergedBranchesShell(2, b) }
func BenchmarkGMGShellBranches3(b *testing.B)  { benchmarkGetMergedBranchesShell(3, b) }
func BenchmarkGMGShellBranches10(b *testing.B) { benchmarkGetMergedBranchesShell(10, b) }
func BenchmarkGMGShellBranches20(b *testing.B) { benchmarkGetMergedBranchesShell(20, b) }
func BenchmarkGMGShellBranches40(b *testing.B) { benchmarkGetMergedBranchesShell(40, b) }

func BenchmarkGMGBranchesCommits1(b *testing.B)  { benchmarkGetMergedBranchesCommits(1, 1, b) }
func BenchmarkGMGBranchesCommits2(b *testing.B)  { benchmarkGetMergedBranchesCommits(2, 2, b) }
func BenchmarkGMGBranchesCommits3(b *testing.B)  { benchmarkGetMergedBranchesCommits(3, 3, b) }
func BenchmarkGMGBranchesCommits10(b *testing.B) { benchmarkGetMergedBranchesCommits(10, 10, b) }
func BenchmarkGMGBranchesCommits20(b *testing.B) { benchmarkGetMergedBranchesCommits(20, 20, b) }
func BenchmarkGMGBranchesCommits40(b *testing.B) { benchmarkGetMergedBranchesCommits(40, 40, b) }

func BenchmarkGMGShellBranchesCommits1(b *testing.B) { benchmarkGetMergedBranchesCommitsShell(1, 1, b) }
func BenchmarkGMGShellBranchesCommits2(b *testing.B) { benchmarkGetMergedBranchesCommitsShell(2, 2, b) }
func BenchmarkGMGShellBranchesCommits3(b *testing.B) { benchmarkGetMergedBranchesCommitsShell(3, 3, b) }
func BenchmarkGMGShellBranchesCommits10(b *testing.B) {
	benchmarkGetMergedBranchesCommitsShell(10, 10, b)
}
func BenchmarkGMGShellBranchesCommits20(b *testing.B) {
	benchmarkGetMergedBranchesCommitsShell(20, 20, b)
}
func BenchmarkGMGShellBranchesCommits40(b *testing.B) {
	benchmarkGetMergedBranchesCommitsShell(40, 40, b)
}

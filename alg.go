// Package alg implements some useful algorithms that are not so trivial to
// implement yet very useful in some situations.
package alg

import (
	"github.com/goutil/num"
)

// Levenshtein computes the minimum distance for converting string a into string
// b, using only Substitution, Deletion and Insertion. This implementation was
// used to solve http://www.spoj.com/problems/EDIST.
func Levenshtein(a, b string) int {
	n := len(a)
	m := len(b)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			neq := 0
			if a[i-1] != b[j-1] {
				neq = 1
			}
			dp[i][j] = dp[i-1][j-1] + neq                 // Substitution
			dp[i][j] = num.MinInt(dp[i][j], dp[i-1][j]+1) // Deletion
			dp[i][j] = num.MinInt(dp[i][j], dp[i][j-1]+1) // Insertion
		}
	}
	return dp[n][m]
}

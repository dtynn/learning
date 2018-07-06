package distinct_subsequences

import "testing"

func TestDistinctSeq(t *testing.T) {
	S := "babgbag"
	T := "bag"

	t.Log(numDistinct(S, T))
}

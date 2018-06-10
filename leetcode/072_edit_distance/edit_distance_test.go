package edit_distance

import (
	"testing"
)

func TestMinDistance(t *testing.T) {
	t.Run("intention > execution", func(t *testing.T) {
		word1 := "intention"
		word2 := "execution"
		t.Log(minDistance(word1, word2))
	})

}

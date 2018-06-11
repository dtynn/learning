package combinations

import (
	"testing"
)

func TestCombaination(t *testing.T) {
	t.Run("3of5", func(t *testing.T) {
		n := 5
		k := 3
		res := combine(n, k)
		for _, line := range res {
			t.Log(line)
		}
	})
}

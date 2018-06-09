package climbing_stairs

import (
	"testing"
)

func TestClimbingStairs(t *testing.T) {
	t.Run("45", func(t *testing.T) {
		t.Log(climbStairs(45))
	})
}

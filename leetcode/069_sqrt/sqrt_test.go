package sqrt

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	t.Run("123001092395803294", func(t *testing.T) {
		x := 123001092395803294
		t.Log(mySqrt(x))
	})
}

package text_justification

import (
	"testing"
)

func TestTextJustification(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		words := []string{"Thisjustificati", "is", "s", "a", "example", "of", "text", "justification."}
		maxWidth := 16
		lines := fullJustify(words, maxWidth)
		for i := range lines {
			t.Logf("%q", lines[i])
		}
	})

	t.Run("3", func(t *testing.T) {
		words := []string{"Listen", "to", "many,", "speak", "to", "a", "few."}
		maxWidth := 6
		lines := fullJustify(words, maxWidth)
		for i := range lines {
			t.Logf("%q", lines[i])
		}
	})

	t.Run("3", func(t *testing.T) {
		words := []string{"Imagination", "is", "more", "important", "than", "knowledge."}
		maxWidth := 14
		lines := fullJustify(words, maxWidth)
		for i := range lines {
			t.Logf("%q", lines[i])
		}
	})

	t.Run("4", func(t *testing.T) {
		words := []string{"My", "momma", "always", "said,", "\"Life", "was", "like", "a", "box", "of", "chocolates.", "You", "never", "know", "what", "you're", "gonna", "get."}
		maxWidth := 12
		lines := fullJustify(words, maxWidth)
		for i := range lines {
			t.Logf("%q", lines[i])
		}
	})
}

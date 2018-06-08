package text_justification

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	lines := make([][]string, 0)
	inline := make([]string, 0)
	inlineWidth := 0

	for {
		if len(words) == 0 {
			break
		}

		word := words[0]
		wordLen := len(word)
		if wordLen > maxWidth {
			return []string{}
		}

		spaceMin := 0
		if len(inline) >= 1 {
			spaceMin = len(inline)
		}

		if inlineWidth+spaceMin+wordLen > maxWidth {
			// if len(lines) > 0 && len(inline) == 1 && canBorrow(lines[len(lines)-1], inline[0], maxWidth) {
			// 	lastLine := lines[len(lines)-1]
			// 	lastWord := lastLine[len(lastLine)-1]
			// 	lines[len(lines)-1] = lastLine[:len(lastLine)-1]
			// 	inline = append([]string{lastWord}, inline...)
			// }

			copied := make([]string, len(inline))
			copy(copied, inline)
			lines = append(lines, copied)
			inline = inline[:0]
			inlineWidth = 0
			continue
		}

		inline = append(inline, word)
		inlineWidth += len(word)
		words = words[1:]
	}

	res := make([]string, 0, len(lines)+1)
	for i := range lines {
		res = append(res, concat(lines[i], maxWidth))
	}

	if len(inline) > 0 {
		lastLine := strings.Join(inline, " ")
		if len(lastLine) < maxWidth {
			lastLine += strings.Repeat(" ", maxWidth-len(lastLine))
		}
		res = append(res, lastLine)
	}

	return res
}

func concat(inline []string, maxWidth int) string {
	if len(inline) == 1 {
		return inline[0] + strings.Repeat(" ", maxWidth-len(inline[0]))
	}

	var inlineWidth int
	for i := range inline {
		inlineWidth += len(inline[i])
	}

	spaceNums := make([]int, len(inline)-1)
	baseNum := (maxWidth - inlineWidth) / len(spaceNums)
	extra := (maxWidth - inlineWidth) % len(spaceNums)
	for i := range spaceNums {
		spaceNums[i] = baseNum
		if i < extra {
			spaceNums[i] = spaceNums[i] + 1
		}
	}

	line := ""
	for i := 0; i < len(inline); i++ {
		line += inline[i]
		if i < len(spaceNums) {
			line += strings.Repeat(" ", spaceNums[i])
		}
	}

	return line
}

func canBorrow(inline []string, word string, maxWidth int) bool {
	if len(inline) <= 2 {
		return false
	}

	borrowWord := inline[len(inline)-1]

	if len(borrowWord)+1+len(word) > maxWidth {
		return false
	}

	return true
}

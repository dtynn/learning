package length_of_last_word

func lengthOfLastWord(s string) int {
	b := []byte(s)
	tail := len(b) - 1
	for tail >= 0 && b[tail] == ' ' {
		tail--
	}

	if tail < 0 {
		return 0
	}

	head := tail - 1
	for head >= 0 && b[head] != ' ' {
		head--
	}

	head++

	return tail - head + 1
}

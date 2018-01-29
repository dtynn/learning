package LetterCombinationOfAPhoneNumber

func letterCombinations(digits string) []string {
	letters := map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}

	var res []string

	for i, b := range []byte(digits) {
		l, ok := letters[b]
		if !ok {
			return []string{}
		}

		if i == 0 {
			res = make([]string, len(l))
			copy(res, l)
			continue
		}

		newres := make([]string, 0, len(res)*len(l))

		for _, origin := range res {
			for _, tail := range l {
				newres = append(newres, origin+tail)
			}
		}

		res = newres
	}

	return res
}

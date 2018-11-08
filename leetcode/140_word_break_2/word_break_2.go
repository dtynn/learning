package word_break_2

import (
	"sort"
)

func wordBreak(s string, wordDict []string) []string {
	memo := map[string][]string{}
	sort.Slice(wordDict, func(i, j int) bool {
		return len(wordDict[j]) < len(wordDict[i])
	})

	return cut(memo, s, wordDict)
}

func cut(memo map[string][]string, s string, dict []string) []string {
	if splits, ok := memo[s]; ok {
		res := make([]string, len(splits))
		copy(res, splits)
		return res
	}

	if s == "" {
		return nil
	}

	possibles := make([]string, 0)
	exists := map[string]struct{}{}

	for i := range dict {
		word := dict[i]
		if len(s) < len(word) {
			continue
		}

		if s == word {
			possibles = append(possibles, word)
			continue
		}

		if s[:len(word)] != word {
			continue
		}

		left, right := s[:len(word)], s[len(word):]
		var lSplits, rSplits []string
		if len(left) > len(right) {
			lSplits = cut(memo, left, dict)
			if len(lSplits) > 0 {
				rSplits = cut(memo, right, dict)
			}

		} else {
			rSplits = cut(memo, right, dict)
			if len(rSplits) > 0 {
				lSplits = cut(memo, left, dict)
			}
		}

		if len(lSplits) > 0 && len(rSplits) > 0 {
			for iL := range lSplits {
				for iR := range rSplits {
					combine := lSplits[iL] + " " + rSplits[iR]
					if _, ok := exists[combine]; !ok {
						exists[combine] = struct{}{}
						possibles = append(possibles, lSplits[iL]+" "+rSplits[iR])
					}
				}
			}
		}
	}

	memo[s] = possibles

	return possibles
}

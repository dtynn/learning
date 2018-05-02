package group_anagrams

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	maxIdx := 0
	idxMap := map[string]int{}
	res := make([][]string, 0)

	for i := range strs {
		bs := []byte(strs[i])
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})

		s := string(bs)
		idx, ok := idxMap[s]
		if !ok {
			idx = maxIdx
			idxMap[s] = idx
			res = append(res, []string{})
			maxIdx++
		}

		res[idx] = append(res[idx], strs[i])
	}

	return res
}

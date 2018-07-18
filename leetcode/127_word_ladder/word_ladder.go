package word_ladder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordset := map[string]struct{}{}
	for _, word := range wordList {
		wordset[word] = struct{}{}
	}

	wordset[beginWord] = struct{}{}

	if _, ok := wordset[endWord]; !ok {
		return 0
	}

	visited := map[string]struct{}{}
	queue := [][]string{
		{endWord},
	}

	for {
		if len(queue) == 0 {
			break
		}

		count := len(queue)
		for n := 0; n < count; n++ {
			curpath := queue[n]
			curword := curpath[0]
			visited[curword] = struct{}{}

			for j := 0; j < len(curword); j++ {
				bs := []byte(curword)
				// 这里受到字符串的拼接方式影响, 结果相差1倍
				for i := 'a'; i <= 'z'; i++ {
					bs[j] = byte(i)
					w := string(bs)
					if _, ok := wordset[w]; ok && w != curword {
						if _, ok := visited[w]; !ok {
							newPath := make([]string, len(curpath)+1)
							copy(newPath[1:], curpath)
							newPath[0] = w

							if w == beginWord {
								return len(curpath) + 1
							}

							queue = append(queue, newPath)
						}
					}
				}
			}
		}
		queue = queue[count:]
	}

	return 0
}

package word_ladder_2

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	res := make([][]string, 0)

	wordset := map[string]struct{}{}
	for _, word := range wordList {
		wordset[word] = struct{}{}
	}

	wordset[beginWord] = struct{}{}

	if _, ok := wordset[endWord]; !ok {
		return res
	}

	visited := map[string]struct{}{}
	queue := [][]string{
		{endWord},
	}

	for {
		if len(queue) == 0 || len(res) > 0 {
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
								res = append(res, newPath)
							}

							if len(res) == 0 {
								queue = append(queue, newPath)
							}
						}
					}
				}
			}
		}
		queue = queue[count:]
	}

	return res
}

// func findLadders(beginWord string, endWord string, wordList []string) [][]string {
// 	res := make([][]string, 0)
// 	visited := map[string]struct{}{}
// 	queue := [][]string{
// 		{beginWord},
// 	}

// 	steps := 1<<63 - 1

// 	wordset := map[string]struct{}{}
// 	for _, word := range wordList {
// 		wordset[word] = struct{}{}
// 	}

// 	if _, ok := wordset[endWord]; !ok {
// 		return res
// 	}

// 	for {
// 		if len(queue) == 0 {
// 			break
// 		}

// 		curpath := queue[0]

// 		if len(curpath) >= steps {
// 			break
// 		}

// 		queue = queue[1:]

// 		curword := curpath[len(curpath)-1]

// 		for i := 'a'; i <= 'z'; i++ {
// 			for j := 0; j < len(curword); j++ {
// 				w := curword[:j] + string([]rune{i}) + curword[j+1:]
// 				if _, ok := wordset[w]; ok && w != curword {
// 					newPath := make([]string, len(curpath)+1)
// 					copy(newPath, curpath)
// 					newPath[len(curpath)] = w

// 					if _, ok := visited[w]; !ok {
// 						queue = append(queue, newPath)
// 					}

// 					if w == endWord {
// 						if len(curpath) < steps {
// 							steps = len(curpath) + 1
// 						}

// 						res = append(res, newPath)
// 					}
// 				}
// 			}

// 			visited[curword] = struct{}{}
// 		}
// 	}

// 	return res
// }

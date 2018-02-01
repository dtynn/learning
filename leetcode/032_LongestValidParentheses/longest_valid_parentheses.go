package LongestValidParentheses

func longestValidParentheses(s string) int {
	longest := 0
	subs := map[int]int{}

	bytes := []byte(s)
	stack := make([]int, 0, len(bytes))
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == '(' {
			stack = append(stack, i)
			continue
		}

		if len(stack) == 0 {
			continue
		}

		// 一组闭合的串
		length := i - stack[len(stack)-1] + 1

		// 之前的连续长度
		length += subs[stack[len(stack)-1]]

		if length > longest {
			longest = length
		}
		subs[i+1] = length

		stack = stack[:len(stack)-1]
	}

	return longest
}

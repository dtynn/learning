package largest_reactange_in_histogram

func largestRectangleArea(heights []int) int {
	largest := 0
	stack := [][2]int{}

	for i := range heights {
		h := heights[i]

		if len(stack) == 0 || h > heights[stack[len(stack)-1][0]] {
			stack = append(stack, [2]int{i, h})
			continue
		}

		for {
			if len(stack) == 0 || heights[stack[len(stack)-1][0]] <= h {
				break
			}

			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			index, area := last[0], last[1]
			if newArea := area + (i-index-1)*heights[index]; newArea > largest {
				largest = newArea
			}
		}

		var left int
		if len(stack) > 0 {
			left = i - stack[len(stack)-1][0]
		} else {
			left = i + 1
		}

		stack = append(stack, [2]int{i, left * h})
	}

	for i := len(stack) - 1; i >= 0; i-- {
		index, area := stack[i][0], stack[i][1]
		if newArea := area + (len(heights)-index-1)*heights[index]; newArea > largest {
			largest = newArea
		}
	}

	return largest
}

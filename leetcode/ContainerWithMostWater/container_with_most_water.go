package ContainerWithMostWater

func maxArea(height []int) int {
	var max int

	i, j := 0, len(height)-1
	for i < j {
		if height[i] > height[j] {
			area := height[j] * (j - i)
			if area > max {
				max = area
			}
			j--
		} else {
			area := height[i] * (j - i)
			if area > max {
				max = area
			}
			i++
		}
	}

	return max
}

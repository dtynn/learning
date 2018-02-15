package TrappingRainWater

func trap(height []int) int {
	var total int

	var left int
	for left < len(height)-1 {
		right := left + 1
		maxIdx := left + 1
		for right < len(height) {
			if height[right] >= height[left] {
				maxIdx = right
				break
			}

			if height[right] > height[maxIdx] {
				maxIdx = right
			}

			right++
		}

		h := height[left]
		if height[maxIdx] < h {
			h = height[maxIdx]
		}

		for i := left + 1; i < maxIdx; i++ {
			total += h - height[i]
		}

		left = maxIdx
	}

	return total
}

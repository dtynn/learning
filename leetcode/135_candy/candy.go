package candy

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 1
	}

	candies := make([]int, len(ratings))

	candies[0] = 1

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
			for prev := i - 1; prev >= 0; prev-- {
				if ratings[prev] <= ratings[prev+1] {
					break
				}

				if candies[prev] <= candies[prev+1] {
					candies[prev] = candies[prev+1] + 1
				}
			}
		}
	}

	total := 0
	for i := 0; i < len(candies); i++ {
		total += candies[i]
	}

	return total
}

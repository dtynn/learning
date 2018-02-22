package JumpGameII

func jump(nums []int) int {
	jcount := 0

	i := 0
	lastMax := 0
	for {
		// 抵达
		if i == len(nums)-1 {
			break
		}

		// 当前能到达的最远位置
		max := nums[i] + i
		if max < len(nums)-1 {
			// 上一轮循环能确定的范围内, 没有点能到达的位置比 i 远, 因此 [i, lastMax] 范围内的点都可以忽略
			next := lastMax + 1
			for j := lastMax + 1 + 1; j <= max; j++ {
				reach := j + nums[j]
				if reach >= len(nums)-1 {
					next = j
					break
				}

				// 如果能到达的位置相同, 则离当前位置越远越好
				if reach >= next+nums[next] {
					next = j
				}
			}

			i = next

		} else {
			i = len(nums) - 1
		}

		lastMax = max

		jcount++
	}

	return jcount
}

package climbing_stairs

func climbStairs(n int) int {
	p := []int{0, 1, 2, 3}

	for i := 4; i <= n; i++ {
		p = append(p, p[i-1]+p[i-2])
	}

	return p[n]
}

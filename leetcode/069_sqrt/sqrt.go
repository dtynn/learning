package sqrt

func mySqrt(x int) int {
	n := 1
	for {
		sqr := n * n
		if sqr == x {
			return n
		}

		if sqr > x {
			return n - 1
		}

		n++
	}
}

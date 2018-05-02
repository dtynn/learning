package pow_x_n

func myPow(x float64, n int) float64 {
	if n < 0 {
		return myPow(1/x, -n)
	}

	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	return myPow(x*x, n/2) * myPow(x, n%2)
}

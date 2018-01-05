package lec01

type Insertion struct {
}

func (Insertion) Sort(in []int) []int {
	if len(in) <= 1 {
		return in
	}

	for i := 1; i < len(in); i++ {
		curr := in[i]
		j := i - 1

		for ; j >= 0 && in[j] > curr; j-- {
			in[j+1] = in[j]
		}

		in[j+1] = curr
	}

	return in
}

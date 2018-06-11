package minimum_window_substring

func minWindow(s string, t string) string {
	if len(s) < len(t) || len(t) == 0 {
		return ""
	}

	bs := []byte(s)
	bt := []byte(t)

	var best []byte
	missed := len(bt)

	target := map[byte]int{}
	for _, b := range bt {
		target[b]++
	}

	counter := map[byte]int{}

	l, r := -1, 0
OUTTER_LOOP:
	for {
		for missed > 0 {
			if r == len(bs) {
				break OUTTER_LOOP
			}

			b := bs[r]
			r++

			required, ok := target[b]
			if !ok {
				continue
			}

			if l == -1 {
				l = r - 1
			}

			counter[b]++
			if counter[b] <= required {
				missed--
			}
		}

		for missed == 0 {
			b := bs[l]
			l++

			required, ok := target[b]
			if !ok {
				continue
			}

			counter[b]--
			if counter[b] < required {
				if r-l < len(best) || len(best) == 0 {
					best = bs[l-1 : r]
				}

				missed++
			}
		}
	}

	return string(best)
}

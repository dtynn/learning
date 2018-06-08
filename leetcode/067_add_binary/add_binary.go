package add_binary

func addBinary(a string, b string) string {
	var res []byte

	ba := []byte(a)
	bb := []byte(b)

	if len(ba) == 0 {
		res = bb
	} else if len(bb) == 0 {
		res = ba
	} else {
		bits := len(a)
		if len(b) > bits {
			bits = len(b)
		}

		res = make([]byte, bits)
		copy(res[bits-len(ba):], ba)
		for i := 0; i < bits-len(ba); i++ {
			res[i] = '0'
		}

		for i := 0; i < len(bb); i++ {
			resIdx := bits - 1 - i
			bIdx := len(bb) - 1 - i
			res[resIdx] = res[resIdx] + (bb[bIdx] - '0')
		}

		up := false
		for i := bits - 1; i >= 0; i-- {
			if res[i] < '2' {
				continue
			}

			res[i] = res[i] - 2
			if i == 0 {
				up = true
				break
			}

			res[i-1] = res[i-1] + 1
		}

		if up {
			res = append([]byte{'1'}, res...)
		}

	}

	for {
		if len(res) == 1 || res[0] != '0' {
			break
		}

		res = res[1:]
	}

	return string(res)
}

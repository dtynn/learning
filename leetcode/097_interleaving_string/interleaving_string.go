package interleaving_string

func isInterleave(s1 string, s2 string, s3 string) bool {
	return isInterleaveBytes([]byte(s1), []byte(s2), []byte(s3))
}

func isInterleaveBytes(bs1, bs2, bs3 []byte) bool {
	if len(bs1)+len(bs2) != len(bs3) {
		return false
	}

	if len(bs1) == 0 {
		return bytesEqual(bs2, bs3)
	}

	if len(bs2) == 0 {
		return bytesEqual(bs1, bs3)
	}

	if bs1[0] == bs3[0] {
		if isInterleaveBytes(bs1[1:], bs2, bs3[1:]) {
			return true
		}
	}

	if bs2[0] == bs3[0] {
		if isInterleaveBytes(bs1, bs2[1:], bs3[1:]) {
			return true
		}
	}

	return false
}

func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

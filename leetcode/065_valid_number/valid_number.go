package valid_number

func isNumber(s string) bool {
	rs := trimSpace([]rune(s))

	if len(rs) == 0 {
		return false
	}

	ei := 0
	for ei < len(rs) {
		if rs[ei] == 'e' {
			break
		}
		ei++
	}

	var base, exp []rune

	if ei == len(rs) {
		base = trimSign(rs)
		return isBase(base)
	}

	base, exp = trimSign(rs[:ei]), trimSign(rs[ei+1:])
	return isBase(base) && isExp(exp)
}

func trimSpace(rs []rune) []rune {
	for {
		if len(rs) == 0 || rs[0] != ' ' {
			break
		}

		rs = rs[1:]
	}

	for {
		l := len(rs)
		if l == 0 || rs[l-1] != ' ' {
			break
		}

		rs = rs[:l-1]
	}

	return rs
}

func trimSign(rs []rune) []rune {
	if len(rs) > 1 && (rs[0] == '+' || rs[0] == '-') {
		return rs[1:]
	}

	return rs
}

func isNumeric(rs []rune) bool {
	if len(rs) == 0 {
		return false
	}

	for i := range rs {
		if rs[i] < '0' || rs[i] > '9' {
			return false
		}
	}

	return true
}

func isExp(rs []rune) bool {
	if len(rs) == 0 {
		return false
	}

	return isNumeric(rs)
}

func isBase(rs []rune) bool {
	if len(rs) == 0 {
		return false
	}

	doti := 0
	for doti < len(rs) {
		if rs[doti] == '.' {
			break
		}
		doti++
	}

	var numi, numf []rune

	if doti == len(rs) {
		numi = rs

	} else {
		numi = rs[:doti]
		numf = rs[doti+1:]
	}

	if len(numi) == 0 {
		return len(numf) != 0 && isNumeric(numf)
	}

	if len(numf) == 0 {
		return isNumeric(numi)
	}

	return isNumeric(numi) && isNumeric(numf)
}

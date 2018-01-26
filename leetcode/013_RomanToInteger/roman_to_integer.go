package RomanToInteger

var value = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	bytes := []byte(s)
	var n int

	for len(bytes) > 0 {
		read := 1
		v := value[bytes[0]]
		if (bytes[0] == 'I' || bytes[0] == 'X' || bytes[0] == 'C') && len(bytes) > 1 {
			if value[bytes[1]] > value[bytes[0]] {
				v = value[bytes[1]] - value[bytes[0]]
				read++
			}
		}

		n += v
		bytes = bytes[read:]
	}

	return n
}

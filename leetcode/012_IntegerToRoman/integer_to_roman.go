package IntegerToRoman

var (
	str    = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
)

func intToRoman(num int) string {
	var i int
	var s string
	for num > 0 {
		if num >= values[i] {
			num -= values[i]
			s += str[i]
			continue
		}

		i++
	}

	return s
}

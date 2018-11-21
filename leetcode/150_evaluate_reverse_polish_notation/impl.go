package impl

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	size := len(tokens)

	if size == 0 || size == 2 {
		panic("invalid token length: " + strconv.Itoa(size))
	}

	if size == 1 {
		i, _ := strconv.Atoi(tokens[0])
		return i
	}

	remains := make([]string, 0, len(tokens))
	for i := 0; i < size-2; i++ {
		if opFn := op(tokens[i+2]); opFn != nil && op(tokens[i]) == nil && op(tokens[i+1]) == nil {
			remains = append(remains, strconv.Itoa(opFn(tokens[i], tokens[i+1])))
			remains = append(remains, tokens[i+3:]...)
			return evalRPN(remains)
		}

		remains = append(remains, tokens[i])
	}

	remains = append(remains, tokens[size-2:]...)

	return -1
}

func op(s string) func(left, right string) int {
	mustInt := func(a string) int {
		i, _ := strconv.Atoi(a)
		return i
	}

	switch s {
	case "+":
		return func(left, right string) int {
			return mustInt(left) + mustInt(right)
		}

	case "-":
		return func(left, right string) int {
			return mustInt(left) - mustInt(right)
		}

	case "*":
		return func(left, right string) int {
			return mustInt(left) * mustInt(right)
		}

	case "/":
		return func(left, right string) int {
			return mustInt(left) / mustInt(right)
		}

	default:
		return nil
	}
}

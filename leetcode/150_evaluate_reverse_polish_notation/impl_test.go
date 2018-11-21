package impl

import "testing"

func TestEvalRPN(t *testing.T) {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	t.Log(evalRPN(tokens))
}

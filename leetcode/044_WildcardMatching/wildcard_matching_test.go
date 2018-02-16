package WildcardMatching

import "testing"

func TestWildcardMatching(t *testing.T) {
	s := "abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb"
	p := "a***ababababbb*bb"
	t.Log(isMatch(s, p))
}

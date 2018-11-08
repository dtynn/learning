package word_break_2

import "testing"

func TestWordBreak2(t *testing.T) {
	cases := []struct {
		s    string
		dict []string
	}{
		{
			"catsanddog",
			[]string{"cat", "cats", "and", "sand", "dog"},
		},
		{
			"pineapplepenapple",
			[]string{"apple", "pen", "applepen", "pine", "pineapple"},
		},
		{
			"catsandog",
			[]string{"cats", "dog", "sand", "and", "cat"},
		},
		{
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
		},
	}

	for i := range cases {
		t.Logf("%#v", wordBreak(cases[i].s, cases[i].dict))
	}
}

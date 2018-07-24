package palindromePartitioning

import "testing"

func Test_minCut(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "babbbbbcb",
			args: args{
				s: "babbbbbcb",
			},
			want: 2,
		},
		{
			name: "babbbbba",
			args: args{
				s: "babbbbba",
			},
			want: 1,
		},
		{
			name: "ab",
			args: args{
				s: "ab",
			},
			want: 1,
		},
		{
			name: "aab",
			args: args{
				s: "aab",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCut(tt.args.s); got != tt.want {
				t.Errorf("minCut() = %v, want %v", got, tt.want)
			}
		})
	}
}

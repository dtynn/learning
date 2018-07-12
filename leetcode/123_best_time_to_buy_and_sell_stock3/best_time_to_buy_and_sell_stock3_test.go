package best_time_to_buy_and_sell_stock3

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[6,1,3,2,4,7]",
			args: args{
				prices: []int{6, 1, 3, 2, 4, 7},
			},
			want: 7,
		},
		// {
		// 	name: "[2, 1, 1, 2, 0 ,1]",
		// 	args: args{
		// 		prices: []int{2, 1, 1, 2, 0, 1},
		// 	},
		// 	want: 2,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

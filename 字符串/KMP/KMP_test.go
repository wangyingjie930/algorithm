package KMP

import "testing"

func Test_getIndexOf(t *testing.T) {
	type args struct {
		subject string
		search  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{subject: "sfsafasefesfsa", search: "asefes"}, want: 5},
		{name: "case1", args: args{subject: "asssdddeadddddadsa", search: "ddddad"}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIndexOf(tt.args.subject, tt.args.search); got != tt.want {
				t.Errorf("getIndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

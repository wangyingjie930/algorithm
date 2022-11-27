package 暴力破解

import (
	"testing"
)

func TestBF(t *testing.T) {
	type args struct {
		needle string
		search string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{needle: "aaaaa", search: "vvv"}, want: -1},
		{name: "case1", args: args{needle: "aaaaa", search: "aaa"}, want: 0},
		{name: "case1", args: args{needle: "asdasdadweaw", search: "dadwe"}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BF(tt.args.needle, tt.args.search); got != tt.want {
				t.Errorf("BF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRK(t *testing.T) {
	type args struct {
		needle string
		search string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{needle: "aaaaa", search: "vvv"}, want: -1},
		{name: "case1", args: args{needle: "aaaaa", search: "aaa"}, want: 0},
		{name: "case1", args: args{needle: "asdasdadweaw", search: "dadwe"}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RK(tt.args.needle, tt.args.search); got != tt.want {
				t.Errorf("RK() = %v, want %v", got, tt.want)
			}
		})
	}
}
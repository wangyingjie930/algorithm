package BM

import (
	"testing"
)

func Test_generateForBadHash(t *testing.T) {
	type args struct {
		search string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{search: "adsadas"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := generateForBadHash(tt.args.search)
			t.Logf("%v", bc)
		})
	}
}

func Test_badCharRule(t *testing.T) {
	type args struct {
		subject string
		search  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{search: "abc", subject: "sadsdabdsbabcsad"}, want: 10},
		{name: "case1", args: args{search: "wang", subject: "sadsdabdsbabcsad"}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := badCharRule(tt.args.subject, tt.args.search); got != tt.want {
				t.Errorf("badCharRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateGoodRule(t *testing.T) {
	type args struct {
		search string
	}
	tests := []struct {
		name  string
		args  args
	}{
		{name: "case1", args: args{search: "cabcab"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := generateGoodRule(tt.args.search)
			t.Logf("%v, %v", got, got1)
		})
	}
}
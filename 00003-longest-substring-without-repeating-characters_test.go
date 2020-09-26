package main

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"abcabcbb", args{"abcabcbb"}, 3},
		{"bbbbb", args{"bbbbb"}, 1},
		{"pwwkew", args{"pwwkew"}, 3},
		{"", args{""}, 0},
		{"aab", args{"aab"}, 2},
		{"dvdf", args{"dvdf"}, 3},
		{" ", args{" "}, 1},
		{"jbpnbwwd", args{"jbpnbwwd"}, 4},
		{"tmmzuxt", args{"tmmzuxt"}, 5},
		{"pwwkewp", args{"pwwkewp"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

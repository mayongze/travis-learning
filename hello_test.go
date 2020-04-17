package main

import "testing"

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{"Hello World!"}, "!dlroW olleH"},
		{"2", args{"你好,世界!"}, "!界世,好你"},
		{"mty", args{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse(%v) = %q, want %q", tt.args, got, tt.want)
			}
		})
	}
}

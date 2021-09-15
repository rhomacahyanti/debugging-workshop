package main

import (
	_ "net/http/pprof"
	"testing"
)

func Test_validateSignal(t *testing.T) {
	type args struct {
		signal []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "good signal",
			args: args{
				signal: []int{1, 0, 1, 0, 0, 1},
			},
			want: "good",
		},
		{
			name: "bad signal",
			args: args{
				signal: []int{0, 0, 1, 0, 0, 0},
			},
			want: "bad",
		},
		{
			name: "unexpected signal",
			args: args{
				signal: []int{0, 0, 1, 0, 0, 7},
			},
			want: "bad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateSignal(tt.args.signal); got != tt.want {
				t.Errorf("validateSignal() = %v, want %v", got, tt.want)
			}
		})
	}
}

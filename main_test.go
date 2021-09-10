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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateSignal(tt.args.signal); got != tt.want {
				t.Errorf("validateSignal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_evaluateSignal(t *testing.T) {
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
				signal: []int{0, 1, 1, 0, 0, 1},
			},
			want: "good",
		},
		{
			name: "bad signal",
			args: args{
				signal: []int{1, 0, 1, 1, 0, 0},
			},
			want: "bad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluateSignal(tt.args.signal); got != tt.want {
				t.Errorf("calculateSignal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_evaluateNode(t *testing.T) {
	type args struct {
		node1 int
		node2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				node1: 1,
				node2: 0,
			},
			want: 0,
		},
		{
			name: "test2",
			args: args{
				node1: 0,
				node2: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluateNode(tt.args.node1, tt.args.node2); got != tt.want {
				t.Errorf("evaluateNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"
)

func TestFibonacci_0(t *testing.T) {
	if Fibonacci(0) != 0 {
		t.Errorf("Expected %v, Actual %v", 0, Fibonacci(0))
	}
}

func TestFibonacci_1(t *testing.T) {
	if Fibonacci(1) != 1 {
		t.Errorf("Expected %v, Actual %v", 1, Fibonacci(1))
	}
}

func TestFibonacci_6(t *testing.T) {
	if Fibonacci(6) != 8 {
		t.Errorf("Expected %v, Actual %v", 6, Fibonacci(8))
	}
}

func TestFibonacci_21(t *testing.T) {
	if Fibonacci(21) != 10946 {
		t.Errorf("Expected %v, Actual %v", 21, Fibonacci(21))
	}
}

func TestFibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "nagative values",
			args: args{n: -1},
			want: -1,
		},
		{
			name: "input too large",
			args: args{n: 21},
			want: 10946,
		},
		{
			name: "input is 1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "input is 2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "input is 3",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "input is 4",
			args: args{n: 4},
			want: 3,
		},
		{
			name: "input is 5",
			args: args{n: 5},
			want: 5,
		},
		{
			name: "input is 6",
			args: args{n: 6},
			want: 8,
		},
		{
			name: "input is 7",
			args: args{n: 7},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.args.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

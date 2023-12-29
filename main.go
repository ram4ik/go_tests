package main

import "fmt"

func main() {
	result := Fibonacci(5)
	fmt.Println(result)
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

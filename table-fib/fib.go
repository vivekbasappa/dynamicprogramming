package main

import "fmt"

func fib(n int) int {
	table := make([]int, n+3, n+3)
	table[1] = 1
	for i := 0 ; i <= n ; i++ {
		table[i+1] += table[i] 
		table[i+2] += table[i] 
	}

	return table[n]
}

func main() {
	fmt.Println(fib(6))
	fmt.Println(fib(7))
	fmt.Println(fib(8))
	fmt.Println(fib(50))
}
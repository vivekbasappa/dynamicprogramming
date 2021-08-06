package main

import "fmt"

// memoization

//fib
func fib(n int, memo map[int]int) int { 
	if val, ok := memo[n]; ok {
		return val 
	}
	if (n <= 2 )  { return 1 }
	memo[n] = fib(n-1, memo) + fib(n-2, memo)
	return memo[n]
}

func main()  {
	cache := make(map[int]int)
	fmt.Printf("fib of 6 : %d\n", fib(6, cache))
	fmt.Printf("fib of 7 : %d\n", fib(7, cache))
	fmt.Printf("fib of 8 : %d\n", fib(8, cache))
	fmt.Printf("fib of 50 : %d\n", fib(50, cache))
}
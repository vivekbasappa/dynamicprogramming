package main

import "github.com/vivekbasappa/dynamicprogramming/util"

// dib simple dib function
func dib(n int) {
	if (n <= 1 ) {
		return
	}
	dib(n-1)
	dib(n-1)
}

func main() {
	defer util.Elapsed("fib")()
	dib(10)
}
package main

import (
	"fmt"
	"github.com/vivekbasappa/dynamicprogramming/util"
)



// fib funciton is to calculate fibonocci serires
func fib(n int) int {
	if (n <= 2) {
		return 1
	}
	return fib(n-1)  + fib(n-2)
}

func main() {
	defer util.Elapsed("fib")()
	value := fib(10)
	fmt.Printf("value %d\n", value) 
}
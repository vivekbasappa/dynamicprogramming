package main

import (
	"context"
	"fmt"

	"github.com/vivekbasappa/dynamicprogramming/util"
)


func fact(n int) int {
	result := 1
	for i:= 2; i < n ; i++ {
			result *= n
	}
	return result
}

type FuncIntInt func(int) int
func memoized(ctx context.Context, fn FuncIntInt) FuncIntInt {
	cache := make(map[int]int)

	return func(input int) int {
		if val, found := cache[input]; found {
			return val
		}
		result := fn(input)
		cache[input] = result
		return result
	}
}

func main() {
	defer util.Elapsed("factorial")()
	factMem := memoized(context.Background(), fact)
	fmt.Println(factMem(5))
	fmt.Println(factMem(5))
	fmt.Println(factMem(5))
	fmt.Println(factMem(5))
	fmt.Println(factMem(5))
	fmt.Println(factMem(5))
	fmt.Println(factMem(5))
}
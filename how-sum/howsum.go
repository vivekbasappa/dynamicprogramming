package main

import "fmt"

// Write a function that takes in a targetSum and an array of numbers as arguments
func howSum(sum int, n []int, memo map[int]*[]int) *[]int {
	if val, ok := memo[sum]; ok { return val }
	if sum == 0 { return &[]int{} }
	if sum <  0 { return nil }
	for _, number := range n {
		difference :=  sum - number
		remainder := howSum(difference, n, memo)
		if remainder != nil {
			val := append(*remainder, number)
			memo[sum] = &val
			return memo[sum]
		}
	}
	memo[sum] = nil
	return nil
}

func main() {
	cache := make(map[int]*[]int)
	fmt.Println(howSum(7, []int{2, 3}, cache))
	cache = make(map[int]*[]int)
	fmt.Println(howSum(7, []int{5, 3, 4, 7}, cache))
	cache = make(map[int]*[]int)
	fmt.Println(howSum(7, []int{2, 4}, cache))
	cache = make(map[int]*[]int)
	fmt.Println(howSum(8, []int{2, 3, 5}, cache))
	cache = make(map[int]*[]int)
	fmt.Println(howSum(300, []int{3, 14}, cache))
}
package main

import "fmt"

// Write a function that takes in a targetSum and an array of numbers as arguments
func howSum(targetSum int, numbers []int, memo map[int]*[]int) *[]int {
	if val, ok := memo[targetSum]; ok { return val }
	if targetSum == 0 { return &[]int{} }
	if targetSum <  0 { return nil }

	var shortestCombination *[]int

	for _, number := range numbers {
		difference :=  targetSum - number
		remainderCombination := howSum(difference, numbers, memo)
		if remainderCombination != nil  {
			combination := append(*remainderCombination, number)
			if shortestCombination == nil || len(combination) < len(*shortestCombination)  {
				shortestCombination = &combination
			}
		}
	}
	memo[targetSum] = shortestCombination
	return shortestCombination
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
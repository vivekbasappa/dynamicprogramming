package main

import "fmt"

func canSum(sum int, n []int, memo map[int]bool) bool {
	if val, ok := memo[sum]; ok { return val }
	if sum == 0 { return true }
	if sum <  0 { return false }
	for _, number := range n {
		difference :=  sum - number
		memo[sum] = canSum(difference, n, memo)
		if memo[sum] {
			return true
		}
	}
	memo[sum] = false
	return false
}

func main() {
	cache := make(map[int]bool)
	fmt.Println(canSum(7, []int{2, 3}, cache))
	cache = make(map[int]bool)
	fmt.Println(canSum(7, []int{5, 3, 4, 7}, cache))
	cache = make(map[int]bool)
	fmt.Println(canSum(7, []int{2, 4}, cache))
	cache = make(map[int]bool)
	fmt.Println(canSum(8, []int{2, 3, 5}, cache))
	cache = make(map[int]bool)
	fmt.Println(canSum(300, []int{3, 14}, cache))
}
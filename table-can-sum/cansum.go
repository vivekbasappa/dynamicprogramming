package main

import (
	"fmt"
)

func canSum(targetSum int, numbers []int) bool {
	table := make([]bool, targetSum+1, targetSum+1)
	table[0] = true
	for index:=0 ; index <= targetSum ; index++ {
		if table[index] == true {
			for _, number := range numbers {
				if index + number < len(table)  {
					table[index+number] = true
				}
			}
		}
	}
	return table[targetSum]
}

func main() {
	fmt.Println(canSum(7, []int{2, 3}))
	fmt.Println(canSum(7, []int{5, 3, 4, 7}))
	fmt.Println(canSum(7, []int{2, 4}))
	fmt.Println(canSum(8, []int{2, 3, 5}))
	fmt.Println(canSum(300, []int{7, 14}))
	fmt.Println(canSum(500, []int{5, 5, 4, 15}))
}
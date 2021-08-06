package main

import (
	"fmt"
)

func gridTravel(m int, n int) int {
	if m == 0 && n == 0 {
		return 0
	}

	table := make([][]int, m+1)
	for i := 0 ; i <= m ; i++ {
		table[i] = make([]int, n+1)
	}
	table[1][1] = 1

	for i := 0 ; i <= m ; i++ {
		for j := 0 ; j <= n; j++ {
			if i+1 <= m {  table[i+1][j] += table[i][j] }
			if j+1 <= n { table[i][j+1] += table[i][j] }
		}
	}
	return table[m][n]
}

func main() {
	fmt.Println(gridTravel(1,1))
	fmt.Println(gridTravel(2,3))
	fmt.Println(gridTravel(3,2))
	fmt.Println(gridTravel(3,3))
	fmt.Println(gridTravel(18,18))

}
package main

import "fmt"

// Grid represents the Grid 
// Memo is used in caching data inbetween
type Grid struct {
	Height int
	Width int
	Memo map[string]int
}

func NewGrid() *Grid {
	 grid := new(Grid)
	 grid.Memo = make(map[string]int)
	 return grid
}

func (g *Grid) CountTravel(m int, n int) int {
	key1 := fmt.Sprintf("%d,%d", m, n)
	key2 := fmt.Sprintf("%d,%d", n, m)
	if val, ok := g.Memo[key1]; ok { return val }
	if val, ok := g.Memo[key2]; ok { return val }
	if m == 1 && n == 1 { return 1 }
	if m == 0 || n == 0 { return 0 }
	g.Memo[key1] = g.CountTravel(m-1, n) + g.CountTravel(m, n-1)
	g.Memo[key1] = g.Memo[key2]	
	return g.Memo[key1]
}

func main() {
	grid := NewGrid()
	fmt.Println(grid.CountTravel(1,1))
	fmt.Println(grid.CountTravel(2,3))
	fmt.Println(grid.CountTravel(3,2))
	fmt.Println(grid.CountTravel(3, 3))
	fmt.Println(grid.CountTravel(18, 18))
}
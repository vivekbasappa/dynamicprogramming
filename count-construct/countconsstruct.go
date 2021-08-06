package main

import (
	"fmt"
	"strings"
)

func countConstruct(target string, wordbank []string, memo map[string]int) int {
	if value, ok := memo[target]; ok { return value }	
	if len(target) == 0 { return 1 } 

	totalCount := 0

	for _, word := range wordbank {
		if strings.HasPrefix(target, word) {
			suffix := strings.SplitN(target, word, 2)
			numWays := countConstruct(suffix[1], wordbank, memo) 
			totalCount += numWays

		}
	}
	memo[target] = totalCount
	return totalCount
}

func main() {
	cache := make(map[string]int)
	fmt.Println(countConstruct("string", []string{"str", "ing", "st", "r", "i", "ng"}, cache))
	cache = make(map[string]int)
	fmt.Println(countConstruct("string", []string{"sr", "ig", "st", "r"}, cache))
	cache = make(map[string]int)
	fmt.Println(countConstruct("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a", "aaaaa", "aaaaaaaa", "aaaaaaaaaaaaaa"}, cache))
}
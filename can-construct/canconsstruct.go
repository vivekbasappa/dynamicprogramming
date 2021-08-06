package main

import (
	"fmt"
	"strings"
)

func canConstruct(target string, wordbank []string, memo map[string]bool) bool {
	if value, ok := memo[target]; ok { return value }
	if len(target) == 0 { return true } 
	
	for _, word := range wordbank {
		if strings.HasPrefix(target, word) {
			suffix := strings.SplitN(target, word, 2)
			if canConstruct(suffix[1], wordbank, memo) {
				memo[target] = true
				return true
			}
		}
	}
	memo[target] = false
	return false
}

func main() {
	cache := make(map[string]bool)
	fmt.Println(canConstruct("string", []string{"str", "ing", "st", "r"}, cache))
	cache = make(map[string]bool)
	fmt.Println(canConstruct("string", []string{"sr", "ig", "st", "r"}, cache))
	cache = make(map[string]bool)
	fmt.Println(canConstruct("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a", "aaaaa", "aaaaaaaa", "aaaaaaaaaaaaaa"}, cache))
}
package main

import (
	"fmt"
	"strings"
)

func allConstruct(target string, wordbank []string)  [][]string {
	if len(target) == 0 { 
		return [][]string{ {""} } 
	} 

	result := [][]string{}

	for _, word := range wordbank {
		if strings.HasPrefix(target, word) {
			suffix := strings.SplitN(target, word, 2)
			suffixWays := allConstruct(suffix[1], wordbank) 
			targetWays := [][]string{}
			for _, numway := range suffixWays {
				numway := append(numway, word)
				targetWays = append(targetWays, numway)
			}
			result  = append(result, targetWays...)
		}
	}
	return result
}

func main() {
	fmt.Println(allConstruct("string", []string{"str", "ing", "st", "r", "i", "ng"}))
	fmt.Println(allConstruct("string", []string{"sr", "ig", "st", "r"}))
	fmt.Println(allConstruct("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a", "aaaaa", "aaaaaaaa", "aaaaaaaaaaaaaa"}))
}
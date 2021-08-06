package util

import (
	"fmt"
	"time"
)

// Elapased this function is used to calculate elapsed time
func Elapsed(what string) func() {
    start := time.Now()
    return func() {
        fmt.Printf("%s took %v\n", what, time.Since(start))
    }
}
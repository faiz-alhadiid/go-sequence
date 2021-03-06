package main

import (
	"fmt"

	"github.com/faiz-alhadiid/go-sequence"
)

func main() {
	result, _ := sequence.
		FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		Filter(func(x int) bool { return x%2 == 0 }).
		Map(func(x int) int { return x * x }).
		Skip(2).
		Reduce(999, func(acc, x int) int {
			if acc < x {
				return acc
			}
			return x
		})
	fmt.Println(result)
}

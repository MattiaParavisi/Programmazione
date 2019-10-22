package main

import (
	"fmt"
	"math"
)

func minmax(s []float64) (min, max float64) {
	min, max = s[0], s[0]
	if len(s) > 1 {
		mi, ma := minmax(s[1:])
		return math.Min(mi, min), math.Max(ma, max)
	}

	return

}

func main() {
	slice := []float64{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	min, max := minmax(slice)
	fmt.Println(min, max)

}

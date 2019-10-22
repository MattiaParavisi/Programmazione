package main

import (
	"fmt"
)

func somma(a *[SIZE]int) {
	for i := 1; i < SIZE; i++ {
		(*a)[i] += (*a)[i-1] // è uguale a dire a[i]+=a[i-1], ma questa è meglio

	}
}

const SIZE = 5

func main() {
	a := [SIZE]int{1, 2, 3, 4, 27}
	fmt.Println("a prima", a)
	somma(&a)
	fmt.Println("a dopo", a)
}

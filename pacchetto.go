package main

import (
	"Palyn/Check"
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	ok := check.IterIsPalin(s)
	if ok {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}
}

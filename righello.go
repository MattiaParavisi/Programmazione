package main

import "fmt"

func tacca(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func righello(n int) {

	if n > 0 {
		righello(n - 1)
		tacca(n)
		righello(n - 1)

	}

}

func main() {

	var n int
	fmt.Scan(&n)
	righello(n)

}

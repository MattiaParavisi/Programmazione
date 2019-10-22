package main

import "fmt"

func aiuto(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("AIUTOH")
	}
}

func main() {
	var a int
	fmt.Println("quante volte vuoi chiedere aiuto?")
	fmt.Scan(&a)
	aiuto(a)

}

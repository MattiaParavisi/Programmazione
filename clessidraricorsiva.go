package main

import "fmt"

func righetta(n int, i int) {
	for i > 0 {
		fmt.Print(" ")
		i--
	}
	for n > 0 {
		fmt.Print("* ")
		n--
	}
	fmt.Println("")
}

func draw(n int, i int) {

	righetta(n, i)
	if n > 1 {
		draw(n-1, i+1)
		righetta(n, i)
	}

}

func main() {
	var n int
	fmt.Println("inserisci la base della clessidra e il livello di indentazione: ")
	fmt.Scan(&n)
	draw(n, 0)
}

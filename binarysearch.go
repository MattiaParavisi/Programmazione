package main

import (
	"fmt"
	"sort"
)

func search(s slice, x int) {

}

func main() {
	slice := []int{1, 2, -2, 3, -5, 6, 7, 0}
	sort.Ints(slice)
	fmt.Println(slice)
	fmt.Print("cerca: ")
	var x int
	fmt.Scan(&x)
	fmt.Println(x)
	fmt.Printf("%d si trova o va inserito nella posizione %d", x, search(slice, x))
}

// ricerca binaria  in una slice s implementata in modo ricorsivo
// restituisce l'indice in cui si trova (la prima occorrenza di) x in s
// o in cui andrebbe inserito x, se non c'Ã¨
// il valore restituito Ã¨ compreso tra 0 e len(s), inclusi.

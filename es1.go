package main

import "fmt"

// definire una funzione Extend che aggiunge un elemento alla fine di una slice di int
// la funzione ritorna una slice con l'aggiunta del nuovo elemento
// definire quindi una funzione Append che, basandosi su Extend,
// aggiunge un numero arbitrario di elementi ad una slice di int (usare ... per denotare #variabile di args)
// la funzioneritorna una slice con l'aggiunta dei nuovi elementi

func Extend(slice []int, i int) []int {
	n := len(slice)
	if n == cap(slice) {
		s := make([]int, n, 2*n+1)
		copy(s, slice)
		slice = s //slice dove copiare,slice da copiare
	}
	slice = slice[0 : n+1]
	slice[n] = i
	return slice

}

func Append(slice []int, i ...int) []int {

	for _, r := range i {
		slice = Extend(slice, r)
	}

	return slice

}

func main() {
	//Extend
	fmt.Println("proviamo extend")
	slice := make([]int, 0, 5)
	for i := 0; i < 10; i++ {
		slice = Extend(slice, i)
		fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
		fmt.Println("indirizzo elemento 0:", &slice[0])
	}

	//Append
	slice = []int{0, 1, 2, 3, 4}
	fmt.Printf("\nprima di append\n%v\n", slice)
	slice = Append(slice, 5, 6, 7, 8)
	fmt.Printf("dopo\n%v", slice)
	slice2 := []int{9, 10, 11, 12, 13, 14}
	slice = Append(slice, slice2...) // i puntini sono necessari! OCIO CHE IL PARAMETRO DI APPEND NON E' DI TIPO SLICE!!! facendo però slice2... funziona
	fmt.Printf("\ndopo secondo append\n%v", slice)

}

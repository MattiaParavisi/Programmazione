package main

import "fmt"

//definire la funzione IsInjective che verifica se una map[string]int rappresenta
//una funzione iniettiva
func IsInjective(m map[string]int) bool {

	m2 := make(map[int]bool)

	for _, c := range m {

		_, ok := m2[c]

		if ok {
			return false
		}

		m2[c] = true

	}
	return true
}

func main() {
	m := map[string]int{"uno": 1, "due": 2, "Uno": 1}
	fmt.Println(m)
	if !IsInjective(m) {
		fmt.Print("non ")
	}
	fmt.Println("iniettiva")
	delete(m, "uno")
	fmt.Println(m)
	if !IsInjective(m) {
		fmt.Print("non ")
	}
	fmt.Print("iniettiva")
}

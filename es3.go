package main

import "fmt"

// definire la funzione ContaSimboli che data una string s ritorna una map
// con le occorrenze dei rune (simboli unicode) in s
func ContaSimboli(s string) map[rune]int {
	occorrenze := make(map[rune]int)
	for _, c := range s {
		occorrenze[c]++
	}
	return occorrenze
}

func main() {
	var s string
	fmt.Scan(&s)
	m := ContaSimboli(s)
	for k, v := range m {
		fmt.Printf("%c%3d\n", k, v)
	}

}

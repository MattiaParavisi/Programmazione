package main

import "fmt"

func mcd(m int, n int) int {
	//considero una coppia del tipo (m,n) e ne calcolo il massimo comun divisore
	if m == 0 {
		return n
	}

	return mcd(n%m, m)

}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(a, b)
	c := mcd(a, b)
	fmt.Println("l'MCD tra i due numeri è:", c)
}

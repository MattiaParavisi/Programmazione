package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	var stringa string
	var res int

	fmt.Scan(&stringa)

	fmt.Println(len(stringa))

	for i := 0; i < len(stringa); i++ {

		if !unicode.IsNumber(rune(stringa[0])) {
			fmt.Printf("Il primo operatore deve essere un numero!")
			return
		}

		if stringa[i] == '+' {
			if stringa[i-2]
			operatore1, _ := strconv.Atoi(string(stringa[i-1]))
			if res == 0 {
				res += operatore1
			}
			operatore2, _ := strconv.Atoi(string(stringa[i+1]))
			res += operatore2
			res := strconv.Itoa(res)
			stringa = res + stringa[i+2:]
			fmt.Println(stringa)
		}

		/*if stringa[i] == '-' {
			operatore1, _ := strconv.Atoi(string(stringa[i-1]))
			if res == 0 {
				res += operatore1
			}
			operatore2, _ := strconv.Atoi(string(stringa[i+1]))
			res -= operatore2
		}*/

		i = 0
	}

	fmt.Println(res)

}

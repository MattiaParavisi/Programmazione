package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func applyOp(res int, acc int, op rune) int {
	if op == '-' {
		return res - acc
	} else if op == '+' {
		return res + acc
	} else {
		return res
	}
}

func main() {
	var expr string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		expr += scanner.Text()
	}
	if strings.HasPrefix(expr, "-") || strings.HasPrefix(expr, "+") {
		expr = "0" + expr
	}

	numeri := make([]int, 0)
	operatori := make([]rune, 0)
	exprAsSlice := []rune(expr)
	for i := 0; i < len(exprAsSlice); i++ {
		if exprAsSlice[i] == '+' || exprAsSlice[i] == '-' {
			operatori = append(operatori, exprAsSlice[i])
		} else if unicode.IsDigit(exprAsSlice[i]) {
			//Raccoglie tutte le cifre fino a un simbolo
			acc := 0
			for i < len(exprAsSlice) && unicode.IsDigit(exprAsSlice[i]) {
				acc = acc*10 + int(exprAsSlice[i]-'0')
				i++
			}
			i--
			numeri = append(numeri, acc)
		}
	}

	//Stampa tutti i numeri e operatori
	fmt.Println(numeri)
	fmt.Println(operatori)

	//Applica gli operatori per tutti i numeri
	res := numeri[0]
	for i := 0; i < len(operatori); i++ {
		if operatori[i] == '+' {
			res += numeri[i+1]
		} else if operatori[i] == '-' {
			res -= numeri[i+1]
		}
	}

	fmt.Println(res)
}

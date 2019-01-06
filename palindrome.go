package main

import "fmt"

func IsPalin(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var s string
	fmt.Scan(&s)
	ok := IsPalin(s)
	if ok {
		fmt.Println("palindroma")
	} else {
		fmt.Println("madunina che bota che gu pres")
	}

}

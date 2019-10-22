//programma che data un altezza disegna un triangolo rovesciato di quell'altezza
//+ + +
// + + 
//  +
package main

import "fmt"

func main(){
	var numero int
	fmt.Println("numero ?")
	fmt.Scan(&numero)
	for i:=0; i<numero; i++{
		
		for j:=0; j<i ; j++{
			fmt.Print(" ")
		}

		for j:=0; j<numero-i; j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}

}
package main

import "fmt"

func main(){
	var numero int
	fmt.Scan(&numero)
	for i:=0; numero!=0; i++{
		for d:=0; d<numero; d++{
			fmt.Print("*")
		}
		numero -= 1
		fmt.Println("")
	}
}
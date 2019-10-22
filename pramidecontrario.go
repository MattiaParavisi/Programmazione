package main

import "fmt"

func main(){
	var rows int
	k:=0
	fmt.Scan(&rows)
	for i:=1; i<rows; i++{
		k=0
		for space:=1; space<=rows-i;space++{
			fmt.Print(" ")
		}
		for k=0; k<=rows; k++{
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}
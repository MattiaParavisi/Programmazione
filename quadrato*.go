package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	for i:=0;i<n;i++{
			for d:=0; d<n; d++ {
				fmt.Print("*")
			}
			fmt.Println("")
	}
}
package main

import "fmt"

func main(){
	var n,pow int
	fmt.Scan(&n)
	var x int
	for {
		fmt.Scan(&x)
		if x<1{
			return
		}
		pow=1
		var i int
		for i=1; i<x; i++{
			pow *= n
			if pow>x{
				fmt.Println("potenza errata")
				break
			}
			if pow==x{
				fmt.Println("potenza esatta")
				break
			}
		}
		if i==x{
			fmt.Println("potenza errata")
		}
	}
}
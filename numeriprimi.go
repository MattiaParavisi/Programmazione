package main

import "fmt"

func main(){
	var numero int
	fmt.Scan(&numero)
	for i:=0; i<numero ;i++{
			var d int
			for d=2; i>d; d++{
				if i%d==0{
					break
				}
			}
		if i==d{
			fmt.Println(i)
		}	
	}
}
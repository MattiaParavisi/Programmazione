package main

import "fmt"

func main(){
	var s string
	fmt.Scan(&s)
	for i:=0; i<len(s) ;i++{
		if s[i]<128{
			
		}else {
			fmt.Println("non ASCII")
			return
		}
	}
	fmt.Println("ASCII")
}
package main

import "fmt"

import "strings"

func main(){
	var stringa string
	fmt.Scan(&stringa)
	s:=strings.ToLower(stringa)
	for c:='a'; c<='z'; c++{
		count:=0
		for _,i:=range s{
			if c==i{
				count++
			}
		}
		if count!=0{
			fmt.Println(string(c),":",count)
		}
	}
	
}
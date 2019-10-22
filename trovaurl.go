package main

import "fmt"

func main(){
	var url string
	fmt.Scan(&url)
	for i:=0; i<len(url); i++{
		if (url[i]=='/')&&(url[i+1]=='/'){
			i+=2 
			for{
				if url[i]=='/'{
					break
				}else {
					fmt.Printf("%c",url[i])
					i++
				}
				
			}
		}
	}
}
package  main

import "fmt"

func main(){
	var s,cifrata string
	fmt.Scan(&s)
	for i:=0; i<len(s); i++{
		c:=s[i]
		if c>='a'&&c<'a'+13||c>='A'&&c<'A'+13{
			c+=13
		}else if c>='a'+13 && c<='z' ||c>='A'+13&&c<='Z'{
			c-=13
		}
		cifrata+=string(c)
	}
	fmt.Println(cifrata)
}
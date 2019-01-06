package main

import "fmt"

func anag(x string) []string {
	if x == "" {
		return []string{""}
	}
	a := anag(x[1:])
	res := []string{}
	mres := make(map[string]bool)
	for _, s := range a {
		for i := 0; i <= len(s); i++ {
			t := s[:i] + string(x[0]) + s[i:]
			if mres[t] {
				continue
			}
			res = append(res, s[:i]+string(x[0])+s[i:])
			mres[t] = true
		}
	}
	return res
}

func main() {
	var s string
	fmt.Scan(&s)
	res := anag(s)
	fmt.Println(res)
}

package main

import (
	"fmt"
	"io"
)

func sequenze(s []string) map[string]int {

	res := make(map[string]int)

	for _, freq := range s {
		res[freq]++
	}

	return res

}

func main() {
	var s []string
	var app string
	var err error

	for err != io.EOF {
		s = append(s, app)
		_, err = fmt.Scan(&app)
	}

	res := sequenze(s)

	fmt.Print(res)

}

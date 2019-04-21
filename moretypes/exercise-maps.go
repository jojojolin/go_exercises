package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	f := strings.Fields(s) //strip string into words array
	m := map[string] int{} //initialise a map with key of type string
	for _, v := range f{ //loop through each word
		m[v] =m[v]+1 //do not need to check if the key exists because update and instert has the same form in Golang
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}


package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z:=1.0
	prev:= 0.0
	for ; z*z <= x && int(z * 100)!=int(prev * 100); z-=(z*z -x)/(3*z){
		fmt.Println(z)
		prev=z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}


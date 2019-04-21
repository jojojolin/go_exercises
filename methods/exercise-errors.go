package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number:%v",float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0{
		return 0, ErrNegativeSqrt(x)
	}
	z:=1.0
	prev:= 0.0
	for ; z*z <= x && int(z * 100)!=int(prev * 100); z-=(z*z -x)/(3*z){
		fmt.Println(z)
		prev=z
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}


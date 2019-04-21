package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a := 0
	b := 0
	f := 0
	return func() int{
		if a == 0 && b == 0{
			b = 1
		} else if a == 0  && b == 1 && f == 0{
			f = 1
		} else {
			f = a+b
			a = b
			b = f
		}
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}


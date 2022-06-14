package main

import "fmt"

// function closure
func returnFunc(x string) func() {
	// sum := 0
	return func() {
		// fmt.Println(sum)
		fmt.Println(x)
	}
}

func main() {
	returnFunc("Hello")()
	x := returnFunc(("ciao"))
	x()
}
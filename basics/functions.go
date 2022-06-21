package main

import "fmt"


func add(a, b int) (map[string] int, int){ // same as a int, b int
	// fmt.Println(a + b)
	// msg := make(map[string] int)
	// msg["ok"] = a + b
	return  map[string] int{"status": 200}, a+b
}

// auto return 
func sub(a, b int) (r1 map[string] int, r2 int) {
	// run right when the function is about to return or at the end of the function
	// useful for something like file closing
	defer fmt.Println("YO I ENDED RIGHTT")
	r1 = map[string] int {"status": 200}
	r2 = a - b
	println("I AM SUBTRACTING YO")
	return 
}

func main () {
	status, sum := add(3, 5)
	fmt.Println(status, sum)

	status, diff := sub(5, 4)
	fmt.Println("Diff", status["status"], diff)
}
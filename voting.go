package main

import "fmt"

func main() {
	var name string
	fmt.Printf("Enter Your Name: ")
	fmt.Scanln(&name)

	fmt.Printf("Enter your age: ")
	var age int8
	fmt.Scanln(&age)

	// The logic is stupid just wanted to try conditionals
	if age > 18 {
		fmt.Printf("You can vote  %9v \n", name)
	} else if age < 18 {
		fmt.Printf("You cannot vote %9v \n", name)
		fmt.Printf("You will have to wait %3d years... \n", 18 - age)
	} else {
		fmt.Println("Register to get your voting id")
	}
}
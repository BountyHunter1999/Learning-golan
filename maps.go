package main 

import "fmt"

func main() {
	fmt.Println("Yellow")
	// array
	var a [3] int = [3]int {2,3,4}
	// fmt.Println(a) => [3 4]

	// a1 := [3]int{4,5,6}
	// // fmt.Println(a1) => [3 4]
	// // slices
	// var b [] int = [] int {3,4}
	// // fmt.Println(b) => [3 4]

	// c := make([]int, 4)
	// c[0] = 1
	// // fmt.Println(c, cap(c), len(c)) => [1 0 0 0] 4 4

	// d := a[:2]
	// d[0] = 1
	// fmt.Println(d, cap(d), len(d))  => [1 3] 3 2


	// ranges
	for i, element := range a {
		println(i, element)
	}
}


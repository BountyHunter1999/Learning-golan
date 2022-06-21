package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

// if we don't make it pointer it will only get the copy
// and changes won't be reflected
// here we haven't made any changes though lol
func (s *Student) avgGrades(grades []int) (avg float32) {
	sum := 0
	for _, elem := range grades {
		sum += elem
	}
	avg = float32(sum / len(grades))
	return
}

func (s *Student) updateAge(newAge int) {
	s.age = newAge
}

func main() {
	s1 := Student{"Shrish", []int{70, 80, 60, 70}, 28}
	fmt.Println(s1.avgGrades(s1.grades))

	fmt.Println("Age before", s1.age)
	s1.updateAge(25)
	fmt.Println("Age before", s1.age)

}

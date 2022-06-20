package main

import (
	"fmt"
	// "time"
)

// import "math"

func main() {
	customers := GetCustomers()

	for x, customer := range customers {
		fmt.Println(x, customer)
	}
}

func getData(customerId int) (customers []string) {

	//
	customers = []string{"Pickachu", "Mewtwo"}

	customers = append(customers, "Ditto")
	customers = append(customers, "Machamp")
	customers = append(customers, "Charmander")

	// for x := 0; x < len(customers); x++ {
	// 	fmt.Println(customers[x])
	// 	time.Sleep(time.Second)
	// }
	for x, customer := range customers {
		fmt.Println(x, customer)
	}
	return customers
}

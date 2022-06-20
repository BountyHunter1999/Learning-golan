package main

type Customer struct {
	FirstName string
	LastName string
	FullName string
}

func GetCustomers()(customers []Customer) {

	pika := Customer{FirstName: "Pika", LastName: "chu", FullName: "Pikachu"}

	customers = append(customers, pika)
	return 
}
package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var andry Customer
	andry.Name = "Andry"
	andry.Address = "Malang"
	andry.Age = 25

	fmt.Println(andry)

	joko := Customer{
		Name:    "Joko",
		Address: "Jakarta",
		Age:     30,
	}

	fmt.Println(joko)
}

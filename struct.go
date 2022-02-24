package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var andry Customer
	andry.Name = "Andry"
	andry.Address = "Malang"
	andry.Age = 25

	andry.sayHi("Joko")

	// fmt.Println(andry)

	// joko := Customer{
	// 	Name:    "Joko",
	// 	Address: "Jakarta",
	// 	Age:     30,
	// }

	// fmt.Println(joko)
}

package main

import "fmt"

func getFullName() (string, string) {
	return "Andry", "Ompusunggu"
}

func main() {
	// fmt.Println(getFullName())

	firstName, _ := getFullName()

	fmt.Println(firstName)
}

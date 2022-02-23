package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{

	// }
	person := map[string]string{
		"name":    "Andry",
		"address": "Malang",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Andry"
	book["ups"] = "Salah"

	delete(book, "ups")

	fmt.Println(book)
}

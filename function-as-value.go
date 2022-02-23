package main

import "fmt"

func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {
	sayGoodbye := getGoodBye

	result := sayGoodbye("Andry")

	fmt.Println(result)
}

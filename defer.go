package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil Function")
}

func runApplication(val int) {
	defer logging()
	fmt.Println("Runc aplication")
	result := 10 / val
	fmt.Println("Result", result)
}

func main() {
	runApplication(0)
}

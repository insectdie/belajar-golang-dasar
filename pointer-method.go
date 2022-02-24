package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	andry := Man{"andry"}
	andry.Married()
	fmt.Println(andry.Name)
}

package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "diubah"
	// fmt.Println(slice1)

	var slice2 = months[8:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "andry")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Andry"
	newSlice[1] = "Ompusunggu"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))

	copy(copySlice, newSlice)
	fmt.Println(copySlice)
	fmt.Println(len(copySlice))
	fmt.Println(cap(copySlice))

}

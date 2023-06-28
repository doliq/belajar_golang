package main

import "fmt"

func main() {
	var months = [12]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	fmt.Println(months[4:7])
	fmt.Println(months[6:12])

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//months[5] = "16 Juni"
	//fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println((slice2))

	var slice3 = append(slice2, "Diwa")
	fmt.Println(slice3)

	newSlice := make([]string, 3, 5)

	newSlice[0] = "Diwa"
	newSlice[1] = "Qotrun"
	newSlice[2] = "Nada"

	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(months))
	copy(copySlice, newSlice)
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
}

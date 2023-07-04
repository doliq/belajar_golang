package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var oliq Customer
	oliq.Name = "Diwa"
	oliq.Address = "Pontianak"
	oliq.Age = 20

	fmt.Println(oliq.Name, oliq.Address)
	fmt.Println(oliq.Address)
	fmt.Println(oliq.Age, "tahun")

	diwa := Customer{
		Name:    "Diwa",
		Address: "Malang",
		Age:     20,
	}
	fmt.Println(diwa)

	arsy := Customer{"Arsy", "Indonesia", 20}
	fmt.Println(arsy)
}

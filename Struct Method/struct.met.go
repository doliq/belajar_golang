package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hi", name, ", My name is ", customer.Name)
}
func main() {
	var oliq Customer
	oliq.Name = "Oliq"
	oliq.Address = "Pontianak"
	oliq.Age = 20

	oliq.sayHi("Diwa")
}

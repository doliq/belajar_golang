package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasname HasName) {
	fmt.Println("Hello", hasname.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}
func main() {
	var oliq Person
	oliq.Name = "Oliq"

	sayHello(oliq)
}

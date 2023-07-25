package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	oliq := Man{"Oliq"}
	oliq.Married()

	fmt.Println(oliq.Name)
}

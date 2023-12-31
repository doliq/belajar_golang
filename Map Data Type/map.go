package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Diwa",
		"address": "Pontianak",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	book := make(map[string]string)
	book["title"] = "Belajar Go-lang"
	book["author"] = "Eko"
	book["ups"] = "wrong"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}

package main

import "fmt"

func main() {
	name := "Oliq"

	if name == "Oliq" {
		fmt.Println("Hello Oliq")
	} else if name == "Diwa" {
		fmt.Println("Hello Diwa")
	} else {
		fmt.Println("Hello Everyone")
	}

	if length := len(name); length > 5 {
		fmt.Println("Your name too long")
	} else {
		fmt.Println("Your name correct")
	}
}

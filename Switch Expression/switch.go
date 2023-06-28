package main

import "fmt"

func main() {
	name := "Arsyad"

	switch name {
	case "Oliq":
		fmt.Println("Hi Oliq")
		fmt.Println("How are you?")
	case "Diwa":
		fmt.Println("Hi Diwa")
		fmt.Println("How are you?")
	case "Arsyad":
		fmt.Println("Hi Arsy")
		fmt.Println("How are you?")
	default:
		fmt.Println("Hi Everyone!")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Name too long")
	// case false:
	// 	fmt.Println("Your name was correct")
	// }

	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Name too long")
	default:
		fmt.Println("Your name was correct")
	}

}

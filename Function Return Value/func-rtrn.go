package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello there"
	} else {
		return "Hello " + name
	}
}

func main() {
	fmt.Println(getHello("Oliq"))
}

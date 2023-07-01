package main

import "fmt"

func getFullName() (string, string, string) {
	return "Diwa", "Arsyad", "Atthoriq"
}

func main() {
	firstName, midName, LastName := getFullName()
	fmt.Println(firstName, midName, LastName)
}

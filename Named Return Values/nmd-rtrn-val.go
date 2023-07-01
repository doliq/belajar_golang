package main

import "fmt"

func getCompleteName() (firstName, midName, lastName string) {
	firstName = "Diwa"
	midName = "Arsyad"
	lastName = "Atthoriq"

	return
}
func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}

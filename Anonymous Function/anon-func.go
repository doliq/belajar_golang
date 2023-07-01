package main

import "fmt"

type Blacklist func(string) bool

func regisUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You're Blocked")
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Oliq"
	}
	regisUser("Diwa", blacklist)
	regisUser("Oliq", blacklist)
}

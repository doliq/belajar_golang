package main

import "fmt"

type Filter func(string) string

func sayHiFilter(name string, filter Filter) {
	filtered := filter(name)
	fmt.Println("Hi", filtered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "***"
	} else if name == "anjing" {
		return "***"
	} else if name == "Babi" {
		return "***"
	} else if name == "babi" {
		return "***"
	} else {
		return name
	}
}

func main() {
	sayHiFilter("Oliq", spamFilter)
	sayHiFilter("Anjing", spamFilter)
	sayHiFilter("Babi", spamFilter)
	sayHiFilter("anjing", spamFilter)
	sayHiFilter("babi", spamFilter)
}

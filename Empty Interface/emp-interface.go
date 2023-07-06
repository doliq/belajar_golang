package main

import "fmt"

func Ups(i int) interface{} {
	// return 1
	// return true
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}

}

func main() {
	var data interface{} = Ups(3)
	fmt.Println(data)
}

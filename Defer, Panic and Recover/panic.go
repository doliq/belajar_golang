package main

import (
	"fmt"
)

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error: ", message)
	}
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("App Error")
	}
	fmt.Println("Done")

}
func main() {
	runApp(false)
	fmt.Println("Diwa")
}

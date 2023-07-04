package main

import "fmt"

func logging() {
	fmt.Println("Done")
}
func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}
}
func main() {
	runApplication()
}

package main

import "fmt"

func main() {
	// {
	// 	tasks := 1

	// 	for tasks <= 10 {
	// 		fmt.Println("Perulangan ke", tasks)
	// 		tasks++
	// 	}

	// }
	for tasks := 1; tasks <= 10; tasks++ {
		fmt.Println("Perulangan ke", tasks)
	}

	slice := []string{
		"Diwa",
		"Arsyad",
		"Atthoriq",
	}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for _, name := range slice {
		// fmt.Println("index", i, "=", name)
		fmt.Println(name)
	}

	person := make(map[string]string)
	person["Name"] = "Diwa"
	person["Title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}

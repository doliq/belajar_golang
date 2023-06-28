package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Diwa"
	names[1] = "Arsyad"
	names[2] = "Atthoriq"

	var d = names[0]
	fmt.Println(d)

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		100,
		95,
		90,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[2])
	fmt.Println(values[1])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var head = [1]string{
		"kepala",
	}
	fmt.Println(head)
}

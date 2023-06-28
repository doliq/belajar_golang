package main

import "fmt"

func main() {

	var name1 = "Oliq"
	var name2 = "Oliq"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 120
	var value2 = 125

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 <= value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}

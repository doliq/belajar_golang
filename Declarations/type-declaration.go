package main

import "fmt"

func main() {
	type NoKTP string

	var noKTPOliq NoKTP = "6112011601030005"
	fmt.Println(noKTPOliq)
}

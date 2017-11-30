package main

import "fmt"

func main() {
	var small, large int
	fmt.Print("What is a small number? ")
	fmt.Scan(&small)
	fmt.Print("What is a large number? ")
	fmt.Scan(&large)
	fmt.Println("The remainder of that division is", large%small)
}

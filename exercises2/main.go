package main

import "fmt"

func main() {
	var input int
	fmt.Print("What number would you like to check for being even? : ")
	fmt.Scan(&input)
	fmt.Println(evencheck(input))
}

func evencheck(z int) bool {
	y := 2
	if z%y == 0 {
		return true
	}
	return false
}

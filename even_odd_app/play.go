package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number to check: ")
	fmt.Scanln(&num)
	res := check(num)
	fmt.Println("Result: ", res)
}

func check(n int) string {
	if n%2 == 0 {
		return "Even Number"
	} else {
		return "Odd"
	}
}

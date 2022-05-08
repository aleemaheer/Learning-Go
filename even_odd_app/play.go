package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number to check: ")
	fmt.Scanln(&num)
	res := check(num)
	fact := factorial(num)
	fmt.Println("Result: ", res)
	fmt.Println("Factorial: ", fact)
}

func check(n int) string {
	if n%2 == 0 {
		return "Even Number"
	} else {
		return "Odd"
	}
}

func factorial(n int) int {
	var i int
	var result int = 1
	for i = n; i > 1; i-- {
		result = result * i
	}
	return result
}

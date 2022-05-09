package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number to check: ")
	fmt.Scanln(&num)
	res := check(num)          // This function will check the number is even or odd
	fact := factorial(num)     // This function will find the factorial of a number
	prime_result := prime(num) // This function will check that the number is even or odd
	fmt.Println("Result: ", res)
	fmt.Println("Factorial: ", fact)
	fmt.Println("IsEven: ", prime_result)
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

func prime(n int) string {
	var i int
	for i = n - 1; i > 1; i-- {
		if n%i == 0 {
			return "Not a Prime Number"
		}
	}
	return "Prime number"
}

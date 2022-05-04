package main

import (
	"example.com/even"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("Input: ")
	log.SetFlags(0)

	check_num, err := even.Even(0)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(check_num)
}

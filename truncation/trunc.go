package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string

	fmt.Println("Enter a floating point number: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatalln("Invalid input.")
		os.Exit(1)
	}

	res := strings.Split(input, ".")

	num, err := strconv.Atoi(res[0])
	if err != nil {
		log.Fatalln("Cannot convert string to integer. Try again.")
		os.Exit(1)
	}

	fmt.Printf("Truncated result is: %d", num)
}

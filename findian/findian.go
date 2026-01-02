package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Enter a string:")

	var reader *bufio.Reader
	reader = bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Invalid input.")
		os.Exit(1)
	}

	// Lower case the string
	input = strings.ToLower(input)
	res, err := regexp.MatchString("^i.*a.*n", input)
	if err != nil {
		log.Fatal("Cannot match string due to error.")
		os.Exit(1)
	}

	if res == true {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

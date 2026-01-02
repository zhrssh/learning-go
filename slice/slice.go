package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func DisplaySlice(slice []int) {
	fmt.Println("Here is the sorted slice you've entered:")

	// Sorts and display the slice
	slices.Sort(slice)
	for idx, val := range slice {
		fmt.Printf("%v\t%v\n", idx+1, val)
	}
}

func main() {
	var slice = make([]int, 0, 3)
	for {
		// Prompts the user to enter integers
		fmt.Println("Enter an integer: ")

		// Read user input
		var reader *bufio.Reader = bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Exit the loop when user enters 'X'
		matched, err := regexp.MatchString("[xX]+", input)
		if err != nil {
			log.Fatal(err)
		}

		if matched {
			DisplaySlice(slice)
			fmt.Println("Exiting program...")
			break
		}

		// Store the integers in a sorted slice
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		slice = append(slice, num)
	}
}

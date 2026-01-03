package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// Initialize reader
	var input string
	var err error
	var reader = bufio.NewReader(os.Stdin)

	// Initialize person map
	var person map[string]string = make(map[string]string)

	// Prompt the user
	fmt.Println("Enter your name:")
	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	person["name"] = strings.TrimSpace(input) // Assign input to name

	fmt.Println("Enter your address:")
	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	person["address"] = strings.TrimSpace(input) // Assign input to address

	fmt.Println("Your JSON object is:")

	// Parse and display bytes
	var bArr []byte
	bArr, err = json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bArr))
}

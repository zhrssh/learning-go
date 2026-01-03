package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	MAX_LENGTH int = 20
)

type Name struct {
	fname string
	lname string
}

func GetNamesFromArray(namesArr []string) ([]Name, error) {
	if len(namesArr) <= 0 {
		return nil, errors.New("There should be at least one name in the array.")
	}

	var Names []Name
	for _, name := range namesArr {
		var trimmed string = strings.TrimSpace(name)
		var split []string = strings.Split(trimmed, " ")

		// Makes sure only those valid splits are appended
		if len(split) <= 1 {
			break
		}

		var newName Name = Name{
			fname: split[0],
			lname: split[1],
		}

		// Add to array
		Names = append(Names, newName)
	}

	return Names, nil
}

func main() {
	// Initialize reader
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var input string
	var err error

	// Prompt the user for the name of the text file
	fmt.Println("Enter filepath to '.txt':")
	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var filepath string = strings.TrimSpace(input)

	// Read names in the file
	fmt.Println("Reading from path: ", filepath)
	var bytes []byte
	bytes, err = os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	// Get the names from the file
	var namesArr []string
	namesArr = strings.Split(string(bytes), "\n")

	// Populate the struct
	var Names []Name
	Names, err = GetNamesFromArray(namesArr)

	// Display the names
	fmt.Println("No,First Name,Last Name")
	for idx, s := range Names {
		fmt.Printf("%d,%s,%s\n", idx, s.fname, s.lname)
	}
}

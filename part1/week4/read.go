package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	var People []Person
	Reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a file name: ")
	FileName, err := Reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	ReadFile, err := os.Open(strings.Trim(FileName, "\n"))
	
	if err != nil {
		panic(err)
	}

	FileScanner := bufio.NewScanner(ReadFile)
	FileScanner.Split(bufio.ScanLines)

	for FileScanner.Scan() {
		Name := strings.Split(FileScanner.Text(), " ")
		People = append(People, Person{fname: Name[0], lname: Name[1]})
	}

	for _, Person := range People {
		fmt.Println(Person.fname, Person.lname)
	}

	ReadFile.Close()
}
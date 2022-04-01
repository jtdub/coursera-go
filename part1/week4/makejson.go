package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"encoding/json"
)
 
func GatherInput(criteria string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your %s: ", criteria)
	input, _ := reader.ReadString('\n')

	return strings.Trim(input, "\n")
}

type Person struct {
	Name string
	Address string
}

func main() {
	j, _ := json.Marshal(Person{Name: GatherInput("name"), Address: GatherInput("address")})
	fmt.Println(string(j))
}
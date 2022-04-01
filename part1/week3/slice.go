package main

import (
	"fmt"
	"strconv"
	"sort"
)

func main() {
	var Loop = true
	var InputValue string
	s := make([]int, 3)

	for Loop {
		fmt.Println("Press 'X' to exit.")
		fmt.Printf("Enter a number to add to a slice: ")
		fmt.Scan(&InputValue)
		
		switch InputValue {
		case "X":
			Loop = false
		default:
			InputValue, _ := strconv.Atoi(InputValue)
			s = append(s, InputValue)
			sort.Ints(s)
			fmt.Printf("%d\n", s)
		}
	}
}
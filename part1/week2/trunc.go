package main

import "fmt"

func main() {
	var floatingNumber float32

	fmt.Printf("Floating Point Number: ")
	fmt.Scan(&floatingNumber)
	fmt.Println(int32(floatingNumber))
}
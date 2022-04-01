package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter a string: ")
	Scanner := bufio.NewScanner(os.Stdin)
	Scanner.Scan()
	SearchString := strings.ToLower(Scanner.Text())

	if strings.HasPrefix(SearchString, "i") && strings.HasSuffix(SearchString, "n") && strings.Contains(SearchString, "a") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
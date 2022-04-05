package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RaiseError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetIntegers() []int {
	sliced_ints := []int{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a sequence of up to ten integers: ")
	input, err := reader.ReadString('\n')
	RaiseError(err)

	for index, item := range strings.Fields(input) {
		if index < 10 {
			item = strings.Trim(item, "\n")
			item, err := strconv.Atoi(item)
			RaiseError(err)
			sliced_ints = append(sliced_ints, item)
		} else {
			break
		}
	}

	return sliced_ints
}

func Swap(sliced_ints *[]int, index int) {
	max_index := len(*sliced_ints) - 1

	// fmt.Println("Max Index Size:", max_index)

	if index < max_index {
		index1 := index + 1
		int0 := (*sliced_ints)[index]
		int1 := (*sliced_ints)[index1]

		// fmt.Println("Performing Swap Operation on index:", index, "and index1:", index1)

		if int0 > int1 {
			(*sliced_ints)[index] = int1
			(*sliced_ints)[index1] = int0

			// fmt.Println(int0, "is larger than", int1)
		}
	}
}

func BubbleSort(sliced_ints *[]int) {
	for i := 0; i < len(*sliced_ints); {
		i++
		for index, _ := range *sliced_ints {
			//fmt.Println("Fetching index:", index)
			Swap(sliced_ints, index)
		}
	}
}

func main() {
	sliced_ints := GetIntegers()
	BubbleSort(&sliced_ints)
	fmt.Println(sliced_ints)
}

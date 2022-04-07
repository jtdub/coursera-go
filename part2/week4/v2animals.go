package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type NewAnimal struct {
	Name    string
	Species Animal
}

type Cow struct {
}

func (c Cow) Eat() string {
	return "grass"
}

func (c Cow) Move() string {
	return "walk"
}

func (c Cow) Speak() string {
	return "moo"
}

type Bird struct {
}

func (b Bird) Eat() string {
	return "worms"
}

func (b Bird) Move() string {
	return "fly"
}

func (b Bird) Speak() string {
	return "peep"
}

type Snake struct {
}

func (s Snake) Eat() string {
	return "mice"
}

func (s Snake) Move() string {
	return "slither"
}

func (s Snake) Speak() string {
	return "hsss"
}

func InputReader() []string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("> ")
	scanner.Scan()
	input := scanner.Text()
	return strings.Split(input, " ")
}

func main() {
	var Animals []NewAnimal

	for true {
		input := InputReader()
		switch input[0] {
		case "newanimal":
			if len(input) == 3 {
				switch input[2] {
				case "cow":
					Animals = append(Animals, NewAnimal{Name: input[1], Species: Cow{}})
					fmt.Println("Created it!")
				case "bird":
					Animals = append(Animals, NewAnimal{Name: input[1], Species: Bird{}})
					fmt.Println("Created it!")
				case "snake":
					Animals = append(Animals, NewAnimal{Name: input[1], Species: Snake{}})
					fmt.Println("Created it!")
				default:
					fmt.Println("Invalid Animal")
					return
				}
			} else {
				fmt.Println("You must create a new animal name and define a species.")
			}
		case "query":
			var animal Animal

			if len(input) == 1 {
				fmt.Println("The requested animal doesn't exist.")
				fmt.Println("Available animals:")
				for _, item := range Animals {
					fmt.Println(item.Name)
				}
			}

			if len(input) == 3 {
				for _, item := range Animals {
					if input[1] == item.Name {
						animal = item.Species
					}
				}

				switch input[2] {
				case "eat":
					fmt.Println(animal.Eat())
				case "move":
					fmt.Println(animal.Move())
				case "speak":
					fmt.Println(animal.Speak())
				default:
					fmt.Println("Invalid verb.")
				}
			}
		default:
			fmt.Println("Invalid command:", input[0])
		}
	}
}

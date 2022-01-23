package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	var userInput string
	var animal Animal
	for {
		fmt.Println(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		userInput = scanner.Text()
		req := strings.Split(userInput, " ")

		switch req[0] {
		case "cow":
			animal = cow
		case "bird":
			animal = bird
		case "snake":
			animal = snake
		default:
			fmt.Println("invalid parameter")
		}

		switch req[1] {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("invalid parameter")
		}

	}

}

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

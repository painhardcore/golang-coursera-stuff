package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat()   { fmt.Println(c.food) }
func (c Cow) Move()  { fmt.Println(c.locomotion) }
func (c Cow) Speak() { fmt.Println(c.noise) }

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (b Bird) Eat()   { fmt.Println(b.food) }
func (b Bird) Move()  { fmt.Println(b.locomotion) }
func (b Bird) Speak() { fmt.Println(b.noise) }

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (s Snake) Eat()   { fmt.Println(s.food) }
func (s Snake) Move()  { fmt.Println(s.locomotion) }
func (s Snake) Speak() { fmt.Println(s.noise) }

func main() {
	var command, param1, param2 string
	var animal, cow, bird, snake Animal
	var animals map[string]Animal
	cow = Cow{"grass", "walk", "moo"}
	bird = Bird{"worms", "fly", "peep"}
	snake = Snake{"mice", "slither", "hsss"}
	animals = map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}
	in := bufio.NewScanner(os.Stdin)
	for {
		command, param1, param2 = "", "", ""
		fmt.Printf(">")
		in.Scan()
		response := strings.ToLower(in.Text())
		words := strings.Fields(response)
		if len(words) == 3 {
			command = (strings.Split(response, " "))[0]
			param1 = (strings.Split(response, " "))[1]
			param2 = (strings.Split(response, " "))[2]
			switch command {
			case "newanimal":
				_, ok := animals[param1]
				if ok {
					fmt.Printf("%s is already present\n", param1)
				} else {
					switch param2 {
					case "cow":
						animals[param1] = cow
					case "bird":
						animals[param1] = bird
					case "snake":
						animals[param1] = snake
					default:
						fmt.Printf("unknown type %s \n", param2)
					}
				}
			case "query":
				_, ok := animals[param1]
				switch ok {
				case true:
					{
						animal = animals[param1]
					}
				case false:
					{
						fmt.Println("unknow animal " + param1)
					}
				}
				switch param2 {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				default:
					fmt.Println("unknow information " + param2)
				}
			}
		} else {
			fmt.Println("invalid request, try again")
		}
	}
}

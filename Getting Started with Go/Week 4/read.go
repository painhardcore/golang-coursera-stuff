package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type person struct {
	lname string
	fname string
}

func main() {
	var path string
	fmt.Println("Enter the name of file:")
	fmt.Scanln(&path)
	crowd := make([]person, 0)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		if len(words) != 2 {
			log.Println("Skipping the line due the incorrect format")
			continue
		}
		crowd = append(crowd, person{fname: words[0], lname: words[1]})
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("In the file:")
	for _, p := range crowd {
		fmt.Printf("%s %s\n", p.fname, p.lname)
	}
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	nameAddress := make(map[string]string, 0)

	fmt.Println("Enter the name")
	scanner.Scan()
	name := scanner.Text()
	nameAddress["name"] = name

	fmt.Println("Enter the address")
	scanner.Scan()
	address := scanner.Text()
	nameAddress["address"] = address

	jsonString, err := json.Marshal(nameAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonString))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter ther word:")
	for scanner.Scan() {
		input := strings.ToLower(scanner.Text())
		if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a") {
			fmt.Println("Found!")
			continue
		}
		fmt.Println("Not Found!")
	}
}

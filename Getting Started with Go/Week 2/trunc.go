package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Printf("cant parse float %s\n", err)
		}
		fmt.Printf("%.f\n", s)
	}
}

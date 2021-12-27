package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	store := make([]int, 0, 3)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter int to add to the slice or \"X\" to exit")
	for scanner.Scan() {
		if strings.ToLower(scanner.Text()) == "x" {
			fmt.Println("Bye!")
			break
		}
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("cant parse int %s\n", err)
		}
		store = append(store, num)
		sort.Ints(store)
		fmt.Printf("Current slice is %v \n", store)
	}
}

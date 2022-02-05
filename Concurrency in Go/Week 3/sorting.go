package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var concurrency int = 4
	var dataSlice = make([]int, 0, concurrency)
	var response string
	var numbers []string
	var size, sizepart int
	var wg sync.WaitGroup

	fmt.Println("Enter integers separeted by space")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	response = in.Text()
	numbers = strings.Split(response, " ")
	size = len(numbers)
	if size < concurrency {
		log.Fatalf("There less than %d integers", concurrency)
	}
	sizepart = size / concurrency

	for _, value := range numbers {
		tmp, err := strconv.Atoi(value)
		if err == nil {
			dataSlice = append(dataSlice, tmp)
		} else {
			log.Fatalf("%s is not an integer", value)
		}
	}
	fmt.Println(dataSlice)
	for i := 0; i < (concurrency - 1); i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, data []int) {
			sort.Ints(data)
			wg.Done()
		}(&wg, dataSlice[(i*sizepart):((i*sizepart)+sizepart)])
	}
	wg.Add(1)
	go func(wg *sync.WaitGroup, data []int) {
		sort.Ints(data)
		wg.Done()
	}(&wg, dataSlice[(sizepart*(concurrency-1)):size])
	// wait sort go routines
	wg.Wait()
	// show sort parts
	for i := 0; i < (concurrency - 1); i++ {
		fmt.Printf("part %d sorted :", i+1)
		fmt.Println(dataSlice[(i * sizepart):((i * sizepart) + sizepart)])
	}
	fmt.Printf("part %d sorted :", concurrency)
	fmt.Println(dataSlice[(sizepart * (concurrency - 1)):size])
	// show merged parts
	fmt.Print("After merge   :")
	fmt.Println(dataSlice)
	// sort sorted parts
	fmt.Print("After sort    :")
	sort.Ints(dataSlice)
	fmt.Println(dataSlice)
}

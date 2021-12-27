package main

import "fmt"

func Swap(arr []int, i int) {
	arr[i], arr[i+1] = arr[i+1], arr[i]
}

func main() {
	var number int
	var arr []int
	for i := 0; i < 10; i++ {
		fmt.Println("Enter a number")
		fmt.Scanln(&number)
		arr = append(arr, number)
	}
	sort := BubbleSort(arr)
	fmt.Println(sort)
}

func BubbleSort(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
	return arr
}

package main

import "fmt"

func findElement(arr [10]int, k int) bool {
	var i int
	for i = 10; i < 10; i++ {
		if arr[i] == k {
			return true
		}
	}
	return false
}

func main() {
	arr := [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	check := findElement(arr, 10)
	fmt.Println(check)
	check = findElement(arr, 9)
	fmt.Println(check)
}

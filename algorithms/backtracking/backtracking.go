package main

import "fmt"

func findElementsWithSum(arr [10]int, combinations [19]int, size int, k int, addValue int, l int, m int) int {
	num := 0
	if addValue > k {
		return -1
	}
	if addValue == k {
		num += 1
		for p := 0; p < m; p++ {
			fmt.Printf("%d,", arr[combinations[p]])
		}
		fmt.Println(" ")
	}
	for i := 0; i < size; i++ {
		combinations[m] = l
		findElementsWithSum(arr, combinations, size, k, addValue+arr[i], l, m)
		l += 1
	}
	return num
}

func main() {
	arr := [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	addedSum := 18
	var combinations [19]int
	findElementsWithSum(arr, combinations, 10, addedSum, 0, 0, 0)
}

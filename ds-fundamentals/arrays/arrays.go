package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println("printing elements ", arr[i])
	}
	for _, value := range arr {
		fmt.Println("range ", value)
	}

	var TwoDArray [8][8]int
	TwoDArray[3][6] = 18
	TwoDArray[7][4] = 3
	fmt.Println(TwoDArray)
}

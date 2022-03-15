package main

import "fmt"

func twiceValue(slice []int) {
	for i, value := range slice {
		slice[i] = 2 * value
	}
}

func main() {

	slice := []int{1, 3, 5, 6}
	slice = append(slice, 8)
	fmt.Println("Capacity", cap(slice))
	fmt.Println("Length", len(slice))

	slice = []int{1, 3, 5, 6}
	twiceValue(slice)
	for i := 0; i < len(slice); i++ {
		fmt.Println("new slice value", slice[i])
	}

	rows := 7
	cols := 9
	twodslices := make([][]int, rows)
	fmt.Println(twodslices)
	for i := range twodslices {
		twodslices[i] = make([]int, cols)
	}
	fmt.Println(twodslices)

	arr := []int{5, 6, 7, 8, 9}
	slice1 := arr[:3]
	fmt.Println("slice1", slice1)
	slice2 := arr[1:5]
	fmt.Println("slice2", slice2)
	slice3 := append(slice2, 12)
	fmt.Println("slice3", slice3)
}

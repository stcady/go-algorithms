package main

import (
	"fmt"
	"log"
)

func powerSeries(a int) (int, int, error) {
	return a * a, a * a * a, nil
}

func PowerSeries(a int) (square int, cube int) {
	square = a * a
	cube = square * a
	return
}

func main() {
	var square int
	var cube int
	var err error
	square, cube, err = powerSeries(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Square: ", square, " Cube: ", cube)
	square, cube = PowerSeries(3)
	fmt.Println("Square: ", square, " Cube: ", cube)
}

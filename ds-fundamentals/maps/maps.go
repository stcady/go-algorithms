package main

import "fmt"

func main() {
	languages := map[int]string{
		3: "English",
		4: "French",
		5: "Spanish",
	}
	products := make(map[int]string)
	products[1] = "chair"
	products[2] = "table"
	for i, value := range languages {
		fmt.Println("language", i, ":", value)
	}
	fmt.Println("product with key 2", products[2])
	delete(products, 1)
	fmt.Println("products", products)
}

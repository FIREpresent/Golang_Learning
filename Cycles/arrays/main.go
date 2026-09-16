package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}
	for _, element := range nums {
		fmt.Printf("Element: %d\n", element)
	}
}

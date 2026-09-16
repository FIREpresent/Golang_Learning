package main

import "fmt"

func main() {
	var money map[string]int = map[string]int{
		"dollars": 1000,
		"euros":   2000,
		"apples":  3,
	}

	// money := map[string]int { ... }

	fmt.Println(money)
	fmt.Println(money["dollars"])

	a, b := money["dollars"]
	fmt.Println(a, b)
}

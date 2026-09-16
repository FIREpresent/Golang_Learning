package main

import "fmt"

func main() {
	var name string
	var age uint8
	fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Println("Hello " + name + "!")
	fmt.Println("How old ar you?")
	fmt.Scan(&age)
	fmt.Println("You are " + fmt.Sprint(age) + " years!")
}

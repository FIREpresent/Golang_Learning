package main

import (
	"fmt"
	"math"
)

// корни квадратного уравнения

func main() {
	var a float64
	var b float64
	var c float64

	fmt.Println("РЕШИ КВАДРАТНОЕ УРАВНЕНИЕ")

	fmt.Println("Введи A:")
	fmt.Scan(&a)

	fmt.Println("Введи B:")
	fmt.Scan(&b)

	fmt.Println("Введи C:")
	fmt.Scan(&c)

	D := b*b - 4*a*c

	if D > 0 {
		var x1 float64
		var x2 float64

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Ваше уравнение имеет 2 корня:")
		fmt.Println(
			"D = " + fmt.Sprint(D),
		)
		fmt.Println(
			"X1 = " + fmt.Sprint(x1) +
				"\nX2 = " + fmt.Sprint(x2),
		)
	} else if D == 0 {
		var x float64

		x = (-b) / (2 * a)

		fmt.Println("Ваше уравнение имеет 1 корень:")
		fmt.Println(
			"D = " + fmt.Sprint(D),
		)
		fmt.Println(
			"X = " + fmt.Sprint(x),
		)

	} else {
		fmt.Println("Ваше уравнение не имеет корней:")
		fmt.Println(
			"D = " + fmt.Sprint(D),
		)
	}

	fmt.Scan(&a)
}

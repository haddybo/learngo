package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

func euler()  {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi) + 1)
}

func triangle()  {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	return c
}

func consts()  {
	const (
		filename = "abc.txt"
		a, b = 3, 4
	)
	var c int
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(filename, c)
}

func main() {
	
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var width int = 15
	var height int = 10

	var perimetr = width*2 + height*2
	var diag float64 = math.Sqrt(float64(width*width + height*height))

	fmt.Println("Периметр прямоугольника = ", perimetr)
	fmt.Println("Длина диагонали прямоугольника = ", diag)
}

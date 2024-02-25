package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a float32 = 10
		b float32 = 20
	)

	s := 2 * math.Pi * a * b
	v := (4 / 3) * math.Pi * a * b * b

	fmt.Printf("Площадь = %1.2f \nОбъем = %1.2f", s, v)
}

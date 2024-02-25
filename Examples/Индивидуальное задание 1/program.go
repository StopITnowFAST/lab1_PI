package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		x1 float64 = 10
		y1 float64 = 0
		z1 float64 = 0
		x2 float64 = 0
		y2 float64 = 0
		z2 float64 = 10
	)

	ans := math.Acos((x1*x2+y1*y2+z1*z2)/
		(math.Sqrt(x1*x1+y1*y1+z1*z1)*
			math.Sqrt(x2*x2+y2*y2+z2*z2))) / math.Pi * 180

	fmt.Printf("Угол между векторами равен %1.2f градуса", ans)
}

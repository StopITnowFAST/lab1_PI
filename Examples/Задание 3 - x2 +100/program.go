package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	x = x*2 + 100

	fmt.Println(x)
}

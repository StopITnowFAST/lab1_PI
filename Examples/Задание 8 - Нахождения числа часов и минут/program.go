package main

import "fmt"

func main() {
	var d, mins, hours int
	fmt.Scan(&d)

	hours = d / 30
	mins = (d - hours*30) * 2

	fmt.Println("It is", hours, "hours", mins, "minutes.")
}

package main

import "fmt"

func main() {
	var n, ans int
	fmt.Scan(&n)
	ans = n % 100 / 10

	fmt.Println(ans)
}

package main

import "fmt"

func main() {
	var n, a, r, i, hasilderet int
	i = 1
	fmt.Scan(&n, &a, &r)
	fmt.Print(0)

	for i <= n {
		hasilderet = a * i * r
		fmt.Print(" + ", hasilderet)
		i++
	}

}

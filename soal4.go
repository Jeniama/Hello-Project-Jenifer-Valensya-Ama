package main

import "fmt"

func main() {
	var n, i int
	var topi, alattulis, tali, pisau bool
	var status bool = true

	fmt.Scan(&n)

	i = 1
	for (i <= n) && status {
		fmt.Scan(&topi, &alattulis, &tali, &pisau)
		status = topi && alattulis && tali && pisau
		i = i + 1
	}
	fmt.Println(status)
}

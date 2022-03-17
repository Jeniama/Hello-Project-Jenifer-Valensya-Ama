package main

import "fmt"

func main() {
	var n, digit, i, jumlah int
	jumlah = 0
	i = 1
	fmt.Scan(&n)
	for i <= n {
		fmt.Scan(&digit)
		for digit < 0 || digit > 9 {
			fmt.Scan(&digit)
		}
		jumlah = jumlah + digit
		i = i + 1
	}
	fmt.Println(jumlah)
}

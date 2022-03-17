package main

import "fmt"

func main() {
	var n, sisapembagian0, sisapembagian1 int
	fmt.Scan(&n)

	var cek bool = true
	sisapembagian1 = n % 10
	sisapembagian0 = sisapembagian1

	n = n / 10

	for n > 0 && cek {
		sisapembagian1 = n % 10
		cek = (sisapembagian1-sisapembagian0) == 1 || (sisapembagian0-sisapembagian1) == 1
		sisapembagian0 = sisapembagian1
		n = n / 10
	}
	fmt.Println(cek)
}

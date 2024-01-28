package main

import "fmt"

func main() {
	var a, b, c, sum int
	for fmt.Scanln(&a); a != 0; a = a / 10 {
		b++
	}

	for b > 0 {
		c = b % 10
		b = b / 10
		sum = sum + c
	}
	fmt.Println(sum)
}

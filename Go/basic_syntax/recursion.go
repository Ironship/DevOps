package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func countDown(n int) {
	if n <= 0 {
		fmt.Println("Done!")
	} else {
		fmt.Println(n)
		countDown(n - 1)
	}
}

func main() {
	fmt.Println(fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))

	countDown(10)
}

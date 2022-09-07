package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	//fmt.Println(fact(7))

	var fib func(n int) int
	times := 0
	fib = func(n int) int {
		if n < 2 {
			return n
		}

		retval := fib(n-1) + fib(n-2)
		times = times + 1
		println(times)
		return retval
	}
	fmt.Println(times)
	fmt.Println(fib(7))
}

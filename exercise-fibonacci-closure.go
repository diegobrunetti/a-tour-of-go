package main

import "fmt"

func fibonacci() func(int) int {
	lastButOne, last := 0, 0

	return func(i int) int {
		if i <= 1 {
			last = i
			return i
		}
		current := lastButOne + last
		lastButOne = last
		last = current

		return current
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

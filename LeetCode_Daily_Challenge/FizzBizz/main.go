package main

import "fmt"

func main() {
	fizzBuzz(15)
}

func fizzBuzz(n int32) {
	// Write your code here
	var i int32
	for i = 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 != 0 && i%5 != 0 {
			fmt.Println(n)
		} else if i%3 == 0 && i%5 != 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println("Buzz")
		}
	}
}

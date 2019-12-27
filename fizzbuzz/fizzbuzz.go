package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
 Count from 1 to a given number n. Printing :
	 - "Fizz" if the number is a multiple of 3
	 - "Buzz" if the number is multiple of 5
	 - "FizzBuzz" if the number is a mutiple of both 3 and 5
*/
func fizzbuzz(n int) {
	for i := 1; i <= n; i++ {
		var fizzbuzzMessage string
		switch {
		case i%15 == 0:
			fizzbuzzMessage = "FizzBuzz"
		case i%3 == 0:
			fizzbuzzMessage = "Fizz"
		case i%5 == 0:
			fizzbuzzMessage = "Buzz"
		}

		fmt.Printf("%d : %s\n ", i, fizzbuzzMessage)
	}
}

func main() {
	n := 50
	if len(os.Args) > 1 {
		max, err := strconv.Atoi(os.Args[1])
		if err == nil {
			n = max
		}
	}
	fizzbuzz(n)
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
  Pick a number n > 1. Apply the following process repeatedly until n becomes 1.
*/
func collatz(n int) int {
	count := 0
	for ; n > 1; count++ {
		var isOdd int = n % 2
		if isOdd == 1 {
			n = n*3 + 1
		} else {
			// n must be even
			n = n / 2
		}
	}
	return count
}

func main() {
	var n int
	var err error
	if len(os.Args) > 1 { // Read the number from the command line
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
	} else { // Read the number interactively
		fmt.Println("Input a number > 1: ")
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if n <= 1 {
		fmt.Println("n must be larger than 1.")
		return
	}
	// := declares, assigns, and infers type in 1 step.
	c := collatz(n)
	fmt.Println(n, "requires", c, "steps to reach 1.")
}

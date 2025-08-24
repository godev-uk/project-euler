package main

import "fmt"

func main() {
	sumSquares := 0
	squareSum := 0
	difference := 0

	for n := 1; n <= 100; n++ {
		sumSquares += n * n
		squareSum += n
	}

	squareSum *= squareSum

	if sumSquares > squareSum {
		difference = sumSquares - squareSum
	} else {
		difference = squareSum - sumSquares
	}

	fmt.Println(difference)
}

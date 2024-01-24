package main

import "fmt"

func main() {
	sum := 0
	sequence := [3]int{1, 1, 2}

	for sequence[2] <= 4000000 {
		// Check if current number is even
		if sequence[2]%2 == 0 {
			sum += sequence[2]
		}

		// Move all numbers back one space,
		// then recalculate current number
		sequence[0] = sequence[1]
		sequence[1] = sequence[2]
		sequence[2] = sequence[0] + sequence[1]
	}

	fmt.Println(sum)
}

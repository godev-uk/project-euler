package main

import "fmt"

func main() {
	const highestFactor = 20
	const lowestFactor = 11
	const requiredFactorCount = 9

	smallestNumber := 0

	for currentNumber := highestFactor; smallestNumber == 0; currentNumber += highestFactor {
		currentFactorCount := 0

		for currentFactor := lowestFactor; currentFactor < highestFactor; currentFactor++ {
			if currentNumber%currentFactor == 0 {
				currentFactorCount++
			} else {
				break
			}
		}

		if currentFactorCount == requiredFactorCount {
			smallestNumber = currentNumber
		}
	}

	fmt.Println(smallestNumber)
}

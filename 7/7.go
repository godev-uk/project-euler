package main

import (
	"fmt"
	"strconv"
)

func getPrimes(maxNumber int) []int {
	const UNKNOWN = 0
	const NOT_PRIME = 1
	const IS_PRIME = 2
	const FIRST_PRIME = 2

	primes := []int{}

	// Use make for the sieve because we know the size and everything should
	// be populated with zero
	sieve := make([]int, maxNumber+1)
	sieve[FIRST_PRIME] = IS_PRIME

	for sieveIndex, currentPrime := FIRST_PRIME, FIRST_PRIME; currentPrime != 0 && sieveIndex <= maxNumber; sieveIndex++ {
		// Assume this prime will be the last
		nextPrime := 0
		step := currentPrime

		// Mark all multiples of current prime as non-prime
		for j := currentPrime + step; j <= maxNumber; j += step {
			sieve[j] = NOT_PRIME
		}

		// Find the next unknown element after the current prime
		// This will be the next prime
		for k := currentPrime + 1; nextPrime == 0 && k <= maxNumber; k++ {
			if sieve[k] == UNKNOWN {
				sieve[k] = IS_PRIME
				nextPrime = k
			}
		}

		// Set current prime to be the next prime
		// If there is no next prime (nextPrime == 0), this will end
		// the for loop as currentPrime != 0 will be false
		currentPrime = nextPrime
	}

	// We now have a sieve with all the numbers flagged as prime/not prime
	// Extract just the primes and return as a slice
	for s := FIRST_PRIME; s <= maxNumber; s++ {
		if sieve[s] == IS_PRIME {
			primes = append(primes, s)
		}
	}

	return primes
}

func findPrimeN(n int, primeDensity int) (int, bool) {
	primes := getPrimes(n * primeDensity)

	if len(primes) >= n {
		return primes[n-1], true
	} else {
		return 0, false
	}
}

func main() {
	const N = 10001

	for primeN, primeDensity, primeFound := 0, 10, false; !primeFound; primeDensity *= 2 {
		fmt.Println("primeDensity: " + strconv.Itoa(primeDensity))
		primeN, primeFound = findPrimeN(N, primeDensity)

		if primeFound {
			fmt.Println("Prime " + strconv.Itoa(N) + " = " + strconv.Itoa(primeN))
		}
	}
}

package main

import (
	"fmt"
	"strconv"
)

/* We shall say that an n-digit number is pandigital if it makes
use of all the digits 1 to n exactly once. For example, 2143 is
a 4-digit pandigital and is also prime.

What is the largest n-digit pandigital prime that exists? */



func PandgitalPrime(n int) int {
	if n > 10 {
		return 0
	}

	for prime := findNextPrime(createNdigitNum(n)); len(strconv.Itoa(prime)) <= n; prime = findNextPrime(prime) {
		if len(strconv.Itoa(prime)) == n && isPandigital(prime) {
			return prime
		}
	}
	return 0
}

func createNdigitNum(n int) int {
	newNum := 9876543210
	for i:=10; i>n; i--{
		newNum = newNum / 10
	}
	return newNum
}

func findNextPrime(N int) int {
	if N <= 1{
		return 2
	}
	// Loop continuously until isPrime returns
	// true for a number greater than n
	for prime := N-1; true; prime-- {
		if isPandigital(prime) && isPrime(prime){
			return prime
		}
	}
	return 0 //unreachable...
}

func isPrime(n int) bool {
	if n <= 1 {return false}
	if n <= 3 {return true}
	if n%2 == 0 || n%3 == 0 {return false}

	for i:=5; i*i<=n; i=i+6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func isPandigital(prime int) bool {
	digits := make([]bool, 10)
	for remainder := 0; prime != 0; prime = prime / 10 {
		remainder = prime % 10
		if digits[remainder] == true {
			return false
		}
		digits[remainder] = true
	}
	return true
}

func main () {
	input := 10
	fmt.Println(PandgitalPrime(input))
}
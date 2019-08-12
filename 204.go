/*
Count the number of prime numbers less than a non-negative number, n.

Example:

Input: 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
 */

package main

func countPrimes(n int) int {
	res := 0
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if !isPrime[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			isPrime[j] = false
		}
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			res++
		}
	}
	return res
}

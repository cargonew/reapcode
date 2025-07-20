package mostFrequent
import "math"
// Given an array of integers, check if there is any integer that appears a prime number of times in the array.
// If there is, return true; otherwise, return false.	
// Example:
// Input: [1, 2, 2, 3, 3, 3, 4]
// Output: true
// Explanation: The number 3 appears 3 times, which is a prime number.
// Input: [1, 1, 2, 2, 3, 3]
// Output: false
// Explanation: No number appears a prime number of times.

func checkPrimeFrequency(nums []int) bool {
    
	freq := make(map[int]int)
	for _,val := range nums{ 	
		freq[val]++
	}

	for _, val := range freq{ 

		if isPrime(val){ 
			return  true
		}
	}
	return false
}

// isPrime checks if a number is prime.

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	// Check for factors from 2 to the square root of nums , its faster that way....
	sqrt := int(math.Sqrt(float64(num)))

	for i := 2; i <= sqrt; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

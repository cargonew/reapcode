package easy

/**
	This program finds the min number off operaions to make array A's sum to be div by int k
	
-How i do this(may not be optimal): 
			+ Sum the whole array 
			+ While the remainder is not zero  , keep subtracting the sum
			+ increment each time you subtract
VARIABLES: 
		counter int : the counter we will return(min ops)
		k int : The divisor/dividor .. or whatever its called 
		nums : The array we will sum
		sum : Where will store the sum
		remainder : store the ramainder

NEW APPROACH : 
		The above one is what i had initially 
		Now i know the answer is r (remainder)
*/
func minOperations(nums []int, k int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	r := sum % k
	return r
}


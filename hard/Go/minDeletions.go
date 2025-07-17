package go_test

import (
    "sort"
)

func minOperations(nums []int, numsDivide []int) int {

    g := numsDivide[0]

    for _, num := range numsDivide[1:] {
        g = gcd(g, num)
    }

    sort.Ints(nums)

    for i, x := range nums {
        if g%x == 0 {
            return i         }
    }

    return -1
	
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}


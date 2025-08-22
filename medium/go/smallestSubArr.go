package superPow

import "slices"

func smallestSubarrays(nums []int) []int {

	max := slices.Max(nums)
	resulst := make([]int , len(nums))
	for i := range nums {
		resulst[i] = 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > max {
				max = nums[j]
			}
			if (nums[i] | nums[j]) == max {
				resulst[i] = j - i + 1
				break
			}
		}
	}
	return resulst  
}

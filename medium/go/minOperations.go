package superPow

// minOperations calculates the minimum number of operations required to make two arrays equal
// by performing operations that add or subtract a given integer k from the elements.
// If it's not possible to make the arrays equal, it returns -1.
// If k is 0, it checks if the arrays are already equal and returns 0
// if they are, or -1 if they are not.
// Time Complexity: O(n), where n is the length of the arrays.
// Space Complexity: O(1), since we are using a constant amount of space.
// Example: minOperations([]int{1, 2, 3}, []int{3, 2, 1}, 2) returns 2,



func minOperations(nums1 []int, nums2 []int, k int) int64 {
	if k == 0 {
		// If k=0, arrays must already be equal
		for i := range nums1 {
			if nums1[i] != nums2[i] {
				return -1
			}
		}
		return 0
	}

	var posDiff, negDiff int64

	for i := 0; i < len(nums1); i++ {
		diff := nums1[i] - nums2[i]

		if diff%k != 0 {
			return -1
		}

		if diff > 0 {
			posDiff += int64(diff / k)
		} else {
			negDiff += int64(-diff / k)
		}
	}

	if posDiff != negDiff {
		return -1
	}

	return posDiff
}


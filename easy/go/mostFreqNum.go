

package mostFrequent

func mostFrequent(nums []int, key int) int {
	frequency := make(map[int]int)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			target := nums[i+1]
			frequency[target]++
		}
	}

	maxCount := 0
	result := -1
	for target, count := range frequency {
		if count > maxCount {
			maxCount = count
			result = target
		}
	}

	return result
}


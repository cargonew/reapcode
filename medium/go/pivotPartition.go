package superPow


func pivotArray(nums []int, pivot int) []int {
	left, middle, right := make([]int, 0), make([]int, 0), make([]int, 0)

	for _, num := range nums {
		if num < pivot {
			left = append(left, num)
		} else if num == pivot {
			middle = append(middle, num)
		} else {
			right = append(right, num)
		}
	}

	result := append(left, middle...)
	result = append(result, right...)
	return result
}


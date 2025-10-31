package mostFrequent

func twoAdjacentIncreasing(nums []int, k int) bool {
    n := len(nums)
    if 2*k > n {
        return false
    }

    isIncreasing := func(start int) bool {
        for i := start; i < start+k-1; i++ {
            if nums[i] >= nums[i+1] {
                return false
            }
        }
        return true
    }

    for i := 0; i <= n-2*k; i++ {
        if isIncreasing(i) && isIncreasing(i+k) {
            return true
        }
    }

    return false
}


package go_test

func maximumLength(nums []int, k int) int {

    n := len(nums)
	
    dp := make([]map[int]int, n)

    maxLen := 1

    for i := 0; i < n; i++ {
        dp[i] = make(map[int]int)

        for j := 0; j < i; j++ {
            mod := (nums[j] + nums[i]) % k

            prevLen := dp[j][mod]
            if prevLen == 0 {
                prevLen = 2
            } else {
                prevLen++ 
            }

            if prevLen > dp[i][mod] {
                dp[i][mod] = prevLen
                if prevLen > maxLen {
                    maxLen = prevLen
                }
            }
        }
    }

    return maxLen
}


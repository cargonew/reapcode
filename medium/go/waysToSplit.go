package superPow

func numberOfGoodSubarraySplits(nums []int) int {
	const mod = 1_000_000_007

	firstOne := -1
	res := 1

	for i, num := range nums {
		if num == 1 {
			if firstOne != -1 {
				zeros := i - firstOne - 1
				res = (res * (zeros + 1)) % mod
			}
			firstOne = i
		}
	}

	if firstOne == -1 {
		return 0
	}

	return res
}




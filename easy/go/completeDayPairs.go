package mostFrequent

func countCompleteDayPairs(hours []int) int {
	freq := make([]int, 24)
	count := 0

	for _, h := range hours {
		mod := h % 24
		complement := (24 - mod) % 24 
		count += freq[complement]
		freq[mod]++
	}

	return count
}


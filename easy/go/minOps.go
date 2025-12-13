package mostFrequent

func minOperations(nums []int, k int) int {

	count := 0

	for _ , val := range nums{ 
		if val < k { 
			count++
		}
	}
	return  count
}

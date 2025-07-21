package mostFrequent

import "slices"

func trimMean(arr []int) float64 {
	slices.Sort(arr)

	cut := len(arr) / 20
	arr = arr[cut : len(arr)-cut]

	sum := 0
	for _, val := range arr {
		sum += val
	}

	return float64(sum) / float64(len(arr))
}


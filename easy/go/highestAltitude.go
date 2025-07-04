
package mostFrequent

func largestAltitude(gain []int) int {
	maxGain := 0
	currGain := 0

	for _, g := range gain {
		currGain += g
		if currGain > maxGain {
			maxGain = currGain
		}
	}

	return maxGain
}


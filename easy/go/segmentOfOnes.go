package mostFrequent


func checkOnesSegment(s string) bool {
	countSegs := 0
	// Count the number of segments of '1's in the string
	// A segment is defined as a contiguous sequence of '1's 
	// that is not preceded by another '1' 
	// For example, in "110011", there are two segments: "11" and "11"
	// In "111", there is one segment: "111"
	// In "101", there are two segments: "1" and "1"
	// In "000", there are no segments of '1's 
	// In "111000", there is one segment: "111"
	for i := 0; i < len(s); i++ {
		if s[i] == '1' && (i == 0 || s[i-1] != '1') {
			countSegs++
		}
	}

	return countSegs <= 1
}


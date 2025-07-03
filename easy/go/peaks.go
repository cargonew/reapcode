package peaks

func findPeaks(mountain []int) []int {
  	
	indices := make([]int, 0)

	for i := 1 ; i < len(mountain)-1 ; i++{
 
		left := mountain[i-1]
		right := mountain[i+1]
		mid := mountain[i]

		if mid > left && mid > right { 
			indices = append(indices, i)
		}
	}

	return  indices
}

package mostFrequent


func threeConsecutiveOdds(arr []int) bool {
    
	for i := 0 ; i < len(arr)-2 ; i++{ 
		
		x := arr[i]
		y := arr[i+1]
		z := arr[i+2]

		if (x % 2 !=0) &&  (y % 2 !=0) && (z  % 2 !=0){ 
			return  true
		}
	}
	return  false
}

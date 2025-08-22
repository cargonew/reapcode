package medium

import "math"

func smallestValue(n int) int {
    
	smallest := 0

	sqrt:= math.Sqrt(n)

	for _,val := range sqrt{ 

		n = if n%val==0{ 

			return val + n/val 
		}
	}

}

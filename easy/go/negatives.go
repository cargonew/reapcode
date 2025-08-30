package mostFrequent

func countNegatives(grid [][]int) int {	
	var negatives int32	
	for _,arr := range grid{ 
		for _,val := range arr{ 
			if val <0{ 
				negatives +=1
			}
		}
	}
	return int(negatives)
}

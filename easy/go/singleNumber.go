
package mostFrequent

func singleNumber(nums []int) int {
    
	results := 0 

	for _,val  := range nums{ 

		results ^= val
	}
	return  results 
}


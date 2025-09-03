package mostFrequent

import "slices" 

func answerQueries(nums []int, queries []int) []int {
   
	slices.Sort(nums)
	results := make([]int , len(queries))

	for i,query := range queries{ 
		for _,num := range nums{
			query -= num
			if query >= 0{
				results[i]++
			}else{ break}
		}
	}
	return results
}

package mostFrequent


func getSneakyNumbers(nums []int) []int {
	seen :=  make(map[int]bool)
	mischiefs := []int{} //two silly numbers
	
	for _,val := range nums { 
		
		if seen[val] { 
			mischiefs = append(mischiefs, val)
		}else { 
			seen[val] = true
		}
	}
	return mischiefs
}

package mostFrequent

func buyChoco(prices []int, money int) int {
    
	min1 , min2 := int(1e9) , int(1e9)
	for _ , p := range prices { 
		if min1 > p { 
			min2 = min1
			min1 = p 

		}else if min2 > p { 
			min2 = p
		}
	}
	total := min1 + min2 

	if total > money { 
		return  money
	} 

	return money - total
}

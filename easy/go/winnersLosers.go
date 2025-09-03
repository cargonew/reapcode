package mostFrequent
func winningPlayerCount(n int, pick [][]int) int {
    max := 0
    keep := make(map[int][]int)
    for _, play := range pick {
        player := play[0]
        ball := play[1]

        keep[player] = append(keep[player], ball)
    }
	for player,itsBalls := range keep{ 
		val := mostFreq(itsBalls)
		if player <val{ 
			max +=1
		}
	}
    return max
}

func mostFreq(arr []int) int{ 	
	maxie := 0 
	freq := make(map[int]int)
	for _,val := range arr{ 
		freq[val]++;
	}
	for _,val := range freq{ 
		if maxie < val{ 
			maxie=val
		}
	}
	return  maxie
}



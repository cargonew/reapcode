package superPow

func findLength(nums1 []int, nums2 []int) int {
	res := 0
	n := maxInt(len(nums1), len(nums2)) 

	for i := range nums1{ 
		val := nums1[i]
		for j := range n { 
			val2 := nums2[j]; 




		}
	}

	return res 
}

func maxInt(a , b int ) int { 
	if a <b { 
	 return a
	}
	return b 
} 

package mostFrequent

func busyStudent(startTime []int, endTime []int, queryTime int) int {
   
	n := len(endTime)
	
	noOfStudents := 0

	for i := range n{ 
		
		if queryTime >=  startTime[i] && queryTime <= endTime[i]{
			noOfStudents++
		} 
	}
    return noOfStudents
}


import (
	"strconv"
	"strings"
)

func convertDateToBinary(date string) string {
	parts := strings.Split(date, "-") 

	binaryParts := make([]string , 0 , len(parts))

	for _,part := range parts { 

		num , err := strconv.Atoi(part)

		if( err != nil){ 
			panic(err)
		}

		binaryNum := strconv.FormatInt(int64(num),2)
		binaryParts = append(binaryParts, binaryNum)

	}

	binaryRep := strings.Join(binaryParts ,"-")

	return binaryRep 
}

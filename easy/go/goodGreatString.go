package mostFrequent

import "unicode"
//This func removes all bad pairs from the string.
// A bad pair is defined as a pair of characters that are the same letter but one is uppercase and the other is lowercase.
// For example, "aA" or "bB" are bad pairs, but "ab" or "AB" are not.

func makeGood(s string) string {


	stack :=[]rune{} // Using a slice as a stack to keep track of characters

	
	for _,ch :=range s{ 

		if len(stack) > 0 && isBadPair(stack[len(stack)-1] , ch){ // Check if the last character in the stack and the current character form a bad pair 	

			stack = stack[:len(stack)-1]
		}else { // append the character to the stack

			stack = append(stack, ch)
		} 
	}
	return  string(stack) 
}

//This function checks if two characters form a bad pair. rturns true if they do else false
func isBadPair( a , b rune) bool { 

	return unicode.ToLower(a) == unicode.ToLower(b) && a !=b
}

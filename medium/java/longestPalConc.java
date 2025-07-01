/**
 * longestPalConc
 */
public class longestPalConc {

    


    public int longestPalindrome(String[] words) {

        StringBuilder str = new StringBuilder();
        
        for ( String word : words){
                
            if( str.append(word).equals( str.append(word).reverse())){

                str.append(word);
            }
        }

        return str.length();        
    }
}


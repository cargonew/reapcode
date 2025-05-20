class countDigits{
 
   public int countDigits(int num) {

        int keep = num ;
        
        int count = 0 ; 

        while ( num >0){
            int val = num % 10 ;

            if( keep % val ==0){
                count++ ;
            }

            num = num / 10 ;

        }
        return count ; 
        
    } 

}

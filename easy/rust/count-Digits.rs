
    pub fn count_digits(mut num: i32) -> i32 {

        let mut count : i32 = 0 ;
        let original = num ;

        while ( num > 0){

            let mut val = num % 10 ; 

            if (original % val ==0 && val != 0  ){
                count+=1;
            } 
            num = num /10; 
        } 
        count 
    }

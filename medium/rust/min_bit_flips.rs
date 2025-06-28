impl Solution {
    pub fn min_flips(a: i32, b: i32, c: i32) -> i32 {
        
        let mut min_flips = 0 ; 

        for i in 0..32{ 

            let abit = ( a>>i) &1;
            let bbit = ( b>>i) &1;
            let cbit = ( c>>i) &1;

            
            if cbit ==1 {

                if abit  ==0 && bbit == 0 {min_flips +=1;}  
            }  

            else { 

                if ( abit ==1 ){min_flips+=1 ;}

                if bbit ==1{ min_flips +=1 ;}
            }
        }
    min_flips
    }
}

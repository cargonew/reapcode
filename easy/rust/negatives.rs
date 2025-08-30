impl Solution {
    pub fn count_negatives(grid: Vec<Vec<i32>>) -> i32 {
        
        let mut negatives : i32 = 0;
        for arr in grid.iter(){ 
            for val in arr.iter(){ 
                if *val  <0 { 
                    negatives +=1 ;
                }
            }
        }
        negatives
    }
    
}

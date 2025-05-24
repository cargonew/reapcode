pub fn maximize_sum(nums: Vec<i32>, k: i32) -> i32 {
        
        let mut max_val = *nums.iter().max().unwrap();
        let mut maxScore = 0 ;

        for i in 0..k{
            maxScore += max_val;
            max_val += 1 ;
        }
        maxScore
    }

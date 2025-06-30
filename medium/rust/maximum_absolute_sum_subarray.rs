impl Solution {
    pub fn max_absolute_sum(nums: Vec<i32>) -> i32 {
            

        let (mut max_sum ,mut  min_sum) = ( 0 , 0);
        let (mut  max_ans ,mut  min_ans) = ( 0 , 0);


        for( num : nums){

            max_sum = ( max_sum+num).max(num); 
            min_sum = ( min_sum+num).min(num);


            max_ans = max_ans.max(max_sum);

            min_ans = min_ans.min(min_sum);
        }
    
    max_ans.max(min_ans.abs())

    }
}

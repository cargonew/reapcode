impl Solution {
    pub fn maximize_sum( mut nums: Vec<i32>, k: i32) -> i32 {

        nums.sort();

        let mut score : i32  = 0 ;

        let n = nums.len()-1;

        for i in 1..=k {

            score +=nums[ n as usize] ; 

            nums[n as usize] = nums[n as usize]+1;
        }
        score
    }
}

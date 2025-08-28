impl Solution {
    pub fn pivot_index(nums: Vec<i32>) -> i32 {
        
        let sum  = nums.iter().sum::<i32>();

        for (i, &num) in nums.iter().enumerate() {
            if sum - num == 2 * nums[..i].iter().sum::<i32>() {
                return i as i32;
            }
        }
        -1
    }
}

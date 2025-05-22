//Literally the first index that meets the requirements is the Solution 
impl Solution {
    pub fn smallest_equal(nums: Vec<i32>) -> i32 {

        for i in 0..nums.len() {

            let val = (i % 10) as i32; //usize to int ( i32)

            if val == nums[i] {
                return i as i32; //again usize to int ( i32)

            }
        }

        -1 //There is no Solution
    }
}


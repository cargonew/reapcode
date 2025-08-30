use std::collections::HashSet;

impl Solution {
    pub fn find_disappeared_numbers(nums: Vec<i32>) -> Vec<i32> {
        let n = nums.len();

        let mut all: HashSet<i32> = (1..=n as i32).collect();
        for &val in nums.iter() {
            all.remove(&val);
        }
        all.into_iter().collect()
    }
}


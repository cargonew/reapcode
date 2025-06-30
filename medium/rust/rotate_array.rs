impl Solution {
    pub fn rotate(nums: &mut Vec<i32>, k: i32) {
        let n = nums.len();
        let k = (k as usize) % n; 
        nums.rotate_right(k);
    }
}


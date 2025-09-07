impl Solution {
    pub fn min_operations(nums: Vec<i32>) -> i32 {
        if nums.len() <= 1 {
            return 0;
        }

        let mut operations = 0;
        let mut prev = nums[0];

        for &num in nums.iter().skip(1) {
            if num <= prev {
                let inc = prev + 1 - num;
                operations += inc;
                prev += 1; 
            } else {
                prev = num;
            }
        }

        operations
    }
}
}

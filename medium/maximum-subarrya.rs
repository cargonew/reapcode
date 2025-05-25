pub fn max_sub_array(nums: Vec<i32>) -> i32 {

        let mut current_sum = nums[0];
        
        let mut max_sum = nums[0];

        for &num in nums.iter().skip(1) {

            current_sum = std::cmp::max(num, current_sum + num);

            max_sum = std::cmp::max(max_sum, current_sum);
        }

        max_sum
    }

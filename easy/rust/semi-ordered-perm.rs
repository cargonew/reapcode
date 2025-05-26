
    pub fn semi_ordered_permutation(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut idx_1 = 0;
        let mut idx_n = 0;

        for (i, &val) in nums.iter().enumerate() {
            if val == 1 {
                idx_1 = i;
            }
            if val == n as i32 {
                idx_n = i;
            }
        }

        if idx_1 < idx_n {
            (idx_1 + (n - 1 - idx_n)) as i32
        } else {
            (idx_1 + (n - 1 - idx_n) - 1) as i32
        }
    }



pub fn maximum_top(nums: Vec<i32>, k: i32) -> i32 {
        let n = nums.len();

        if k == 0 {
            return *nums.get(0).unwrap_or(&-1);
        }

        if n == 1 && k % 2 == 1 {
            return -1;
        }

        if k as usize > n {
            return *nums.iter().max().unwrap();
        }

        let mut my_max = std::i32::MIN;

        for i in 0..std::cmp::min(k - 1, n as i32) {
            my_max = std::cmp::max(my_max, nums[i as usize]);
        }

        if (k as usize) < n {
            my_max = std::cmp::max(my_max, nums[k as usize]);
        }

        my_max
    }

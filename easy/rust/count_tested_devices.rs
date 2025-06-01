impl Solution {
    pub fn count_tested_devices(mut battery_percentages: Vec<i32>) -> i32 {
        let mut count = 0;
        let n = battery_percentages.len();

        for i in 0..n {
            if battery_percentages[i] > 0 {
                count += 1;
                for j in i+1..n {
                    battery_percentages[j] = (battery_percentages[j] - 1).max(0);
                }
            }
        }

        count
    }
}


impl Solution {
    pub fn count_time(time: String) -> i32 {
        


        let mut count = 0;

    for hour in 0..24 {
        for minute in 0..60 {
            let candidate = format!("{:02}:{:02}", hour, minute);
            if time.chars().zip(candidate.chars()).all(|(t, c)| t == '?' || t == c) {
                count += 1;
            }
        }
    }

    count
    }
}

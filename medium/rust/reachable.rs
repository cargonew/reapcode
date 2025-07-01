impl Solution {
    pub fn is_reachable_at_time(sx: i32, sy: i32, fx: i32, fy: i32, t: i32) -> bool {
        let dx = (fx - sx).abs();
        let dy = (fy - sy).abs();
        let steps_needed = dx.max(dy);

        if dx == 0 && dy == 0 && t == 1 {
            return false;
        }

        steps_needed <= t
    }
}


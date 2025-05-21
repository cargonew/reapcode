pub fn tribonacci(n: i32) -> i32 {
        let n = n as usize;
        let mut vals = vec![0, 1, 1];
        if n < 3 {
            return vals[n];
        }
        for i in 3..=n {
            let next = vals[i - 1] + vals[i - 2] + vals[i - 3];
            vals.push(next);
        }
        vals[n]
    }

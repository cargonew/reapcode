impl Solution {
    pub fn super_pow(a: i32, b: Vec<i32>) -> i32 {
        fn mod_pow(x: i32, n: i32, m: i32) -> i32 {
            let mut res = 1;
            let mut base = x % m;
            let mut exp = n;

            while exp > 0 {
                if exp % 2 == 1 {
                    res = (res * base) % m;
                }
                base = (base * base) % m;
                exp /= 2;
            }

            res
        }

        const MOD: i32 = 1337;
        let mut result = 1;

        for &digit in &b {
            result = mod_pow(result, 10, MOD) * mod_pow(a, digit, MOD) % MOD;
        }

        result
    }
}


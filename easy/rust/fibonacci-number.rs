impl Solution {
    pub fn fib(n: i32) -> i32 {
        if n == 0 {
            return 0;
        }
        if n == 1 {
            return 1;
        }

        let mut store = Vec::with_capacity((n + 1) as usize);
        store.push(0);
        store.push(1);

        for i in 2..=n {
            let next = store[(i - 1) as usize] + store[(i - 2) as usize];
            store.push(next);
        }

        store[n as usize]
    }
}


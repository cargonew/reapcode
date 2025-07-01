impl Solution {
    pub fn k_weakest_rows(mat: Vec<Vec<i32>>, k: i32) -> Vec<i32> {
        let mut store = vec![];

        for (i, row) in mat.iter().enumerate() {
            let soldiers = count_ones(row);
            store.push((soldiers, i as i32)); 
        }

        store.sort(); 

        store.iter()
            .take(k as usize)
            .map(|&(_, index)| index)
            .collect()
    }
}

pub fn count_ones(vect: &Vec<i32>) -> i32 {
    let mut count = 0;

    for &val in vect {
        if val == 1 {
            count += 1;
        }
    }

    count
}






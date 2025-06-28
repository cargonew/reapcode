impl Solution {
    pub fn minimum_boxes(apple: Vec<i32>, mut capacity: Vec<i32>) -> i32 {
     
        let mut sum : i32 = apple.iter().sum();

        capacity.sort(); 

        let mut x = capacity.len();
        for i in x-1..0{ 
            sum = sum - capacity[i as usize];
            if sum < 0 { return  i-x ;}
        } 

    }
}

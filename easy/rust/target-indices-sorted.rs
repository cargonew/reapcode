pub fn target_indices(mut nums: Vec<i32>, target: i32) -> Vec<i32> {

        let mut indices = Vec::new();

        nums.sort(); 
        for i in 0..nums.len(){

            if ( nums[i]==target){

                indices.push ( i as i32);
            }
        }
        indices
    }

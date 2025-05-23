
    pub fn min_number_of_hours(mut initial_energy: i32, mut initial_experience: i32, energy: Vec<i32>, experience: Vec<i32>) -> i32 {
        

        let mut total_hours : i32 = 0 ; 

        for i in 0..energy.len(){

            if (initial_energy <= energy[i]){

                let mut needed = energy[i] - initial_energy + 1;

                total_hours +=needed ; 
                initial_energy += needed ;
            } 
            if ( initial_experience <= experience[i]){

                let needed = experience[i] - initial_experience +1 ; 

                total_hours += needed ; 

                initial_experience +=needed ;
            } 

            initial_experience += experience[i];
            initial_energy -= energy[i];
        }
        total_hours
    }

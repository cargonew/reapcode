 pub fn get_longest_subsequence(words: Vec<String>, groups: Vec<i32>) -> Vec<String> {

        //This is basic two pointers amd stuff nothing much 
        //Check if the adjecent elemnts in groups are not the same then add    the corresponding words element at j since i is trailing.

        let mut substring = Vec::new();

        substring.push ( words[0].clone());

        for i in 0 ..words.len()-1{ 
            let j = i + 1; 
            if groups[i] != groups[j] {
                substring.push(words[j].clone());
            }
        }
        substring
    }

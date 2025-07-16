package mostFrequent


func commonFactors(a int, b int) int {
    count := 0 

    for i := 1; i <= 1000 ; i++{ 

        if a%i==0 && b%i==0{ 
            count++
        }
    }
    return count
}

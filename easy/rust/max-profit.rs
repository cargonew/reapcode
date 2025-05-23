 //This one was a little bit tricky but here ...
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let mut min_price = i32::MAX; //this is the initial min price
        let mut max_profit = 0; //This is the intial max_profit

        for price in prices {

            if price < min_price { //Checks if current price is less than the current minimum price .. if it is ..we make it the new minimum price

                min_price = price; //here 
            } else if price - min_price > max_profit { //if its not we continue checking if "today's" profit is bigger than what we already have.
                max_profit = price - min_price; //if it is then its the new max ..we move.
            }
        }
        max_profit //return the max profit... and thats it.
    }

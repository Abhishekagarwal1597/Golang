// you are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

// Function to calculate the maximum profit that can be obtained from buying and selling a stock
func maxProfit(prices []int) int {
    // If the length of the prices array is less than or equal to 1, return 0
    if len(prices) <= 1 {
        return 0
    }
    // Initialize the minimum price to be the first price in the array
    minPrice := prices[0]
    // Initialize the maximum profit to be 0
    maxProfit := 0
    // Iterate over the prices array starting from the second element
    for i := 1; i < len(prices); i++ {
        // If the current price is less than the minimum price, update the minimum price
        if prices[i] < minPrice {
            minPrice = prices[i]
        } else if prices[i] - minPrice > maxProfit {
            // If the difference between the current price and the minimum price is greater than the maximum profit,
            // update the maximum profit
            maxProfit = prices[i] - minPrice
        }
    }
    // Return the maximum profit
    return maxProfit
}

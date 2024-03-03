# To find the number of subarrays whose sum equals a given target k, we can use a hashmap to store the cumulative sum frequencies.
func subarraySum(nums []int, k int) int {
    hashMap := make(map[int]int)
	  var prefixSum, count int
	  for i:=0;i<len(nums);i++{
		    prefixSum += nums[i]
        
        //checking if prefix sum is equal to required sum
    		if prefixSum == k {
    			count++
    		}
        
        //check if prefix sum already exist for required sum
        //add all the previouse hashmap sum
    		if _, ok := hashMap[prefixSum-k]; ok{
    			count += hashMap[prefixSum-k]
    		}
        
        //update the prefixsum count
        hashMap[prefixSum]++
        
	  }
	  return count
}

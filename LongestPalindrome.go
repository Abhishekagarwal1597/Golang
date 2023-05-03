// longestPalindrome takes a string s and returns the length of the longest palindrome that can be built using its characters.
// This function uses a hash table to count the frequency of each character in s. It then loops through the hash table and counts the number of even frequency characters.
// If there is at least one character with an odd frequency count, it adds one to the count. This is because a palindrome can have at most one character with an odd frequency count.
func longestPalindrome(s string) int {
    // Initialize an empty hash table to count the frequency of each character in s
    freq := make(map[rune]int)
    for _, c := range s {
        freq[c]++
    }
    
    // Loop through the hash table and count the number of even frequency characters
    hasOdd := false
    count := 0
    for _, f := range freq {
        if f%2 == 0 {
            count += f
        } else {
            count += f - 1
            hasOdd = true
        }
    }
    
    // Add one to the count if there is at least one character with an odd frequency count
    if hasOdd {
        count++
    }
    
    // Return the count, which represents the length of the longest palindrome that can be built using the characters in s
    return count
}

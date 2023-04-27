//check if string s is a subsequence of a sting t or not
func isSubsequence(s string, t string) bool {
    i, j := 0, 0
    //use for loop to iterate through it 
    for i < len(s) && j < len(t) {
        //if we find a match increse the s index
        if s[i] == t[j] {
            i++
        }
        // t index will always increse
        j++
    }
  //return based on s index count
    return i == len(s)
}

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    //creating map for storing mapping
    s2t := make(map[byte]byte)
    t2s := make(map[byte]byte)
    
    for i := 0; i < len(s); i++ {
        sChar := s[i]
        tChar := t[i]
        
        //check if mapping is already present for existing char or not
        if s2t[sChar] == 0 && t2s[tChar] == 0 {
            //add the mapping if not present
            s2t[sChar] = tChar
            t2s[tChar] = sChar
        } else if s2t[sChar] != tChar || t2s[tChar] != sChar {
            //check if mapping fail for any char return false
            return false
        }
    }
    
    return true
}

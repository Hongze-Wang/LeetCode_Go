func lengthOfLongestSubstring(s string) int {
    bits := make(map[byte]int)
    start, res := 0, 0
    
    for i, _ := range s {
        if start < bits[s[i]] {
            start = bits[s[i]]
        }
        
        if res < i+1-start {
            res = i+1-start               
        }
        
        bits[s[i]] = i+1
    }
    
    return res
}

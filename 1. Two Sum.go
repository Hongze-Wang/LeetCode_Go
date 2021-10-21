// core idea

func twoSum(nums []int, target int) []int {
    seen := map[int]int{}
    res := []int{}
    need := 0
	
    for i, val := range nums {
    	need = target - val
	v, ok := seen[need]

	if ok {
		res = append(res, i, v)
		break
	}

	seen[val] = i
	
    }
	return res
}

// optimized version, more go style
func twoSum(nums []int, target int) []int {
    seen := map[int]int{}
    for i, num := range nums {
	need := target - num
	if _, ok := seen[need]; ok {
	    return []int{seen[need], i}
	}
	seen[num] = i
    }
    return []int{}
}

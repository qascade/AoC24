package missionChief

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
func checkSafety(levels []int) bool {
	isPositive := false
	if len(levels) <= 1 {
		return true
	}

	if levels[1] - levels[0] > 0 {
		isPositive = true
	} 

	for i := 0; i < len(levels) - 1; i++ {
		diff := levels[i+1] - levels[i]
		if isPositive {
			if diff < 0 {
				return false
			}
			if diff < 1 || diff > 3 {
				return false
			}
		} else {
			// levels should be monotonically decreasing
			if diff > 0 {
				return false
			}
			if (-1 * diff) < 1 || (-1 * diff) > 3 {
				return false
			}
		}
	}
	return true
}

// This function returns the number of safe reports
func Day2Part1(reports [][]int) int {
	count := 0
	for _, levels := range reports {
		isSafe := checkSafety(levels)
		if isSafe {
			count += 1
		}
	}
	return count
}

// func checkSafetyWithDampner(levels []int) bool {
// 	if len(levels) <= 1 {
// 		return true
// 	}

// 	for end < len(levels) {
		
// 	}


// 	return true
// }


// func Day2Part2(reports [][]int) int {
// 	// Return Safe if by removing one level the number starts satisfying the above conditions.
// 	// We can try using two pointer/Sliding window approach where a window will not cover
// 	// more than three levels at a time. 
// 	count := 0
// 	for _, levels := range reports {
// 		isSafe := checkSafetyWithDampner(levels)
// 		if isSafe {
// 			count += 1
// 		}
// 	}
// 	return count
// }
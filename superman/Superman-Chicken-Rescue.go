package superman

// return max number of chicken can carry
func supermanRescue(inputLine1, positions []int) (maxCarriedChicken int) {
	//validate both chicken and superman's length for carry the roof and positions is not empty
	if len(inputLine1) == 0 || len(positions) == 0 {
		return 0
	}
	start := 0
	n := inputLine1[0] //Number of chickens
	k := inputLine1[1] //Length of the roof
	// Iterate through each position with the end pointer
	for end := 0; end < n; end++ {

		// Shrink the window from the start if the difference between
		// positions[end] and positions[start] is greater than or equal to k
		for positions[end]-positions[start] >= k {
			start++
		}
		// Update the maximum number of chickens that can be protected
		if end-start+1 > maxCarriedChicken {
			maxCarriedChicken = end - start + 1
		}
	}
	return maxCarriedChicken
}

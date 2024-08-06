package baby

import (
	"strings"
)

const MaxTwoCharInput = 2
const (
	Shoot   = 'S'
)

func babyRevenge(input *string) bool {
	trimmedInput := strings.TrimSpace(*input) // trime all space to make sure that input is ready to process next.
	// validating incoming input is greater than 0 length for futher process.
	if len(trimmedInput) > 0 {

		newInput := strings.ToUpper(trimmedInput) // to support that input are lower case
		if newInput[0] != Shoot{
			return false
		}
		if isValidInput(newInput) {
			return processBaby(&newInput)
		}
	}
	return false
}

// processBaby means; to monitor baby's behavior which is counting the event.
func processBaby(input *string) bool {
	mapBaby := make(map[string]int, MaxTwoCharInput) // mapBaby contains two char which is 'R' and 'S' and have capacity 2 ,based on 2 characters.
	//counting mapBaby["R"] and mapBaby["S"].
	for _, v := range *input {
		mapBaby[string(v)]++
	}
	return mapBaby["S"] < mapBaby["R"] // if mapBaby["S"] is less than mapBaby["R"] indicate that is a good boy.
}

// to validate that in coming input is contain only "S" and "R", if not return false
func isValidInput(s string) bool {
	for _, ch := range s {
		if ch != 'S' && ch != 'R' {
			return false
		}
	}
	return true
}

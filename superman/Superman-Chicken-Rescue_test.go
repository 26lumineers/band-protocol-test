package superman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SupermanChicken struct {
	ScenarioName string
	InputLine1 []int
	Positions  []int
	Expected   int
}

func TestSupermanChickenRescue(t *testing.T) {
	testCases := []SupermanChicken{
		{
			ScenarioName :"Test #1",
			InputLine1: []int{5, 5},
			Positions:  []int{2, 5, 10, 12, 15},
			Expected:   2,
		},
		{
			ScenarioName :"Test #2",
			InputLine1: []int{6, 10},
			Positions:  []int{1, 11, 30, 34, 35, 37},
			Expected:   4,
		},
		{
			ScenarioName :"Empty both chicken, and length of the roof Superman can carry",
			InputLine1: []int{},
			Positions: []int{3, 8, 15, 21, 24, 35, 42, 50, 55, 60},
			Expected:   0,
		},
		{
			ScenarioName :"all empty, input and positions",
			InputLine1: []int{},
			Positions: []int{},
			Expected:   0,
		},
		{
			ScenarioName :"positions is empty",
			InputLine1: []int{1,2},
			Positions: []int{},
			Expected:   0,
		},
		{
			ScenarioName :"Superman can't save anyone",
			InputLine1: []int{5,0},
			Positions: []int{3, 8, 15, 21, 24, 35, 42, 50, 55, 60},
			Expected:   0,
		},
		{
			ScenarioName :"No chickens to save",
			InputLine1: []int{0,5},
			Positions: []int{3, 8, 15, 21, 24, 35, 42, 50, 55, 60},
			Expected:   0,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.ScenarioName, func(t *testing.T) {
			result := supermanRescue(tt.InputLine1, tt.Positions)
			assert.Equal(t,tt.Expected,result,"they should be equal, we got %d, but we want %d",result,tt.Expected)
		})
	}
}

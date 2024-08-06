package baby

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type BabyRevenge struct{
	ScenarioName string
	Name string
	Input string
	Expected bool
}
func TestBabyRevenge(t *testing.T) {
	testcases := []BabyRevenge{
		{
			ScenarioName :"Test #1",
			Name:     "Good boy",
			Input:    "SRSSRRR",
			Expected: true,
		},
		{
			ScenarioName :"Test #2",
			Name:     "Bad boy",
			Input:    "RSSRR",
			Expected: false,
		},
		{
			ScenarioName :"Test #3",
			Name:     "Bad boy",
			Input:    "SSSRRRRS",
			Expected: false,
		},
		{
			ScenarioName :"Test #4",
			Name:     "Bad boy",
			Input:    "SRRSSR",
			Expected: false,
		},
		{
			ScenarioName :"Test #5",
			Name:     "Good boy",
			Input:    "SSRSRRR",
			Expected: true,
		},
		{
			ScenarioName :"Charactors are lower case",
			Name:     "Good boy",
			Input:    "ssrsrrr",
			Expected: true,
		},
		{
			ScenarioName :"Invalid input, charactor is not S and R",
			Name:     "Bad boy",
			Input:    "SSRSRRRz",
			Expected: false,
		},
		{
			ScenarioName :"Invalid input, empty input",
			Name:     "Bad boy",
			Input:    "",
			Expected: false,
		},
		{
			ScenarioName :"Invalid input, charactor is a number",
			Name:     "Bad boy",
			Input:    "123",
			Expected: false,
		},
		

	}

	for _, tt := range testcases {
		t.Run(tt.ScenarioName,func(t *testing.T) {
			result := babyRevenge(&tt.Input)
			if tt.Expected != result {
				assert.Equal(t,tt.Expected,result,"they should be equal, we got %v, but we want %v",result,tt.Expected)
			}
		})
	}
}
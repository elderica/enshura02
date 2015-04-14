package enshura02

import (
	"testing"
	"strings"
)

var testdata = []struct {
	input []int
	expected [7]int
}{
	{
		input: []int{
			30000, 20000, 0, 50000,     0, 100000, 500000,
			30000, 20000, 0, 20000, 15000, 600000, 450000,
		},
		expected: TotalPerDay{
			60000, 40000, 0, 70000, 15000, 700000, 950000,
		},
	},
	{
		input: []int{
			1003, 2302, 421, 32124, 3, 0, 3214,
		},
		expected: TotalPerDay{
			1003, 2302, 421, 32124, 3, 0, 3214,
		},
	},
}

func TestCount(t *testing.T) {
	for i, testdatum := range testdata {
		if actual := Count(testdatum.input); actual != testdatum.expected {
			t.Errorf(
				"No. %d test failed.\n" +
				"expected: %v\n" +
				"actual:   %v\n",
				i, testdatum.expected, actual)
		}
	}
}

var testinput = []string{
`14
30000
20000
0
50000
0
100000
500000
30000
20000
0
20000
15000
600000
450000
`,
`7
1003
2302
421
32124
3
0
3214
`}

func Run(inp string) error {
	r := strings.NewReader(inp)
	err := Main(r)
	if err != nil {
		return err
	}
	return nil
}

func ExampleF_Test01() {
	Run(testinput[0])
	// Output: 
	// 60000
	// 40000
	// 0
	// 70000
	// 15000
	// 700000
	// 950000
}

func ExampleF_Test02() {
	Run(testinput[1])
	// Output:
	// 1003
	// 2302
	// 421
	// 32124
	// 3
	// 0
	// 3214
}

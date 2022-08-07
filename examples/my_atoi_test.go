package examples

import (
	"testing"
)

// Melakukan pengujian dengan penerapan collection values
func TestMyAtoiSubtest(t *testing.T) {
	var tests = []struct {
		Input string
		Want  int
	}{
		{
			Input: "",
			Want:  0,
		},
		{
			Input: "     -46",
			Want:  -46,
		},
		{
			Input: "+46",
			Want:  46,
		},
		{
			Input: "     4567/1234",
			Want:  4567,
		},
		{
			Input: "89876Hello Word",
			Want:  89876,
		},
		{
			Input: "1234567891032",
			Want:  2147483647,
		},
		{
			Input: "-1234567891032",
			Want:  -2147483648,
		},
		{
			Input: "1024",
			Want:  1024,
		},
		{
			Input: "-1024",
			Want:  -1024,
		},
	}

	for _, test := range tests {
		if result := MyAtoi(test.Input); result != test.Want {
			t.Fatalf("Testing with input value %s, expected %d but result %d", test.Input, test.Want, result)
		}
	}

	t.Log("Testing function MyAtoi finished")
}

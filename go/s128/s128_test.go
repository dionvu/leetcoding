package s128

import "testing"

func TestLongestConsec(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
	}

	for _, test := range tests {
		res := longestConsecutive(test.input)

		if res != test.expected {
			t.Fatalf("Recieved [%d], wanted [%d]", res, test.expected)
		}
	}
}

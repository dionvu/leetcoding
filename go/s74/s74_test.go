package s74

import "testing"

func Test74(t *testing.T) {
	tests := []struct {
		input    [][]int
		target   int
		expected bool
	}{
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
	}

	for _, test := range tests {
		res := searchMatrix(test.input, test.target)

		if res != test.expected {
			t.Fatalf("Recieved [%t], wanted [%t]", res, test.expected)
		}
	}
}

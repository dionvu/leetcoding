package s11

import "testing"

func Test11(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
		nth      int
	}{
		{[]int{2, 3, 4, 5, 18, 17, 6}, 17, 2},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49, 1},
	}

	for _, test := range tests {
		res := maxArea(test.input)

		if res != test.expected {
			t.Fatalf("Test #[%d], Recieved [%d], wanted [%d]", test.nth, res, test.expected)
		}
	}
}

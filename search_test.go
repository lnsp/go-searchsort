package searchsort

import "testing"

func TestLinearSearch(t *testing.T) {
	cases := []struct {
		Values  []int
		Subject int
		Exists  bool
	}{
		{[]int{1, 2, 3, 4, 5}, 2, true},
		{[]int{1, 3, 5, 2, 4}, 2, true},
		{[]int{1, 2, 3, 4, 5}, 0, false},
		{[]int{}, 0, false},
	}

	for _, c := range cases {
		r := LinearSearch(c.Values, c.Subject)
		if r != c.Exists {
			t.Error("Expected", c.Exists, "got", r)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		Values  []int
		Subject int
		Exists  bool
	}{
		{[]int{1, 2, 3, 4, 5}, 2, true},
		{[]int{1, 2, 3, 4, 5}, 0, false},
		{[]int{1, 2}, 2, true},
		{[]int{1}, 1, true},
		{[]int{}, 0, false},
	}

	for _, c := range cases {
		r := BinarySearch(c.Values, c.Subject)
		if r != c.Exists {
			t.Error("Expected", c.Exists, "got", r)
		}
	}
}

func TestHashSearch(t *testing.T) {
	var hash HashFunction = func(subject int, size int) int {
		return subject % size
	}

	cases := []struct {
		Values  []int
		Subject int
		Exists  bool
	}{
		{[]int{0, 1, 2, 3, 4}, 0, true},
		{[]int{0, 1, 2, 3, 4}, 5, false},
		{[]int{5, 6, 0, 3, 9}, 6, true},
		{[]int{5, 6, 0, 3, 9}, 0, false},
		{[]int{1}, 1, true},
		{[]int{}, 0, false},
	}

	for _, c := range cases {
		r := HashSearch(hash, c.Values, c.Subject)
		if r != c.Exists {
			t.Error("Expected", c.Exists, "got", r, "for", c.Values)
		}
	}
}

package searchsort

import "testing"
import "math/rand"

func isSorted(x []int) bool {
	l := len(x)
	if l > 0 {
		v := x[0]
		for i := 1; i < l; i++ {
			if x[i] < v {
				return false
			}
			v = x[i]
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		values := make([]int, rand.Intn(1000))
		for j := range values {
			values[j] = rand.Intn(1000)
		}
		BubbleSort(values)
		if !isSorted(values) {
			t.Error(values, "is not sorted")
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	values := make([]int, 1000)
	for j := range values {
		values[j] = 1000
	}
	for i := 0; i < b.N; i++ {
		dest := make([]int, len(values))
		copy(dest, values)
		BubbleSort(dest)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
    values := make([]int, 1000)
	for j := range values {
		values[j] = 1000
	}
	for i := 0; i < b.N; i++ {
		dest := make([]int, len(values))
		copy(dest, values)
		InsertionSort(dest)
	}
}

func TestInsertionSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		values := make([]int, rand.Intn(10000))
		for j := range values {
			values[j] = rand.Intn(10000)
		}
		InsertionSort(values)
		if !isSorted(values) {
			t.Error(values, "is not sorted")
		}
	}
}

func TestSelectionSort(t *testing.T) {
    for i := 0; i < 1000; i++ {
		values := make([]int, rand.Intn(10000))
		for j := range values {
			values[j] = rand.Intn(10000)
		}
		SelectionSort(values)
		if !isSorted(values) {
			t.Error(values, "is not sorted")
		}
	}
}
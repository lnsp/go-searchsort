package searchsort

func BubbleSort(a []int) {
	for n := len(a); n > 0; n-- {
		for i := 0; i < n-1; i++ {
			if a[i+1] < a[i] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
}

func InsertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		ins := a[i]
		j := i
		for j > 0 && a[j-1] > ins {
			a[j] = a[j-1]
			j--
		}
		a[j] = ins
	}
}

func SelectionSort(a []int) {
	n := len(a)
	for l := 0; l < n; l++ {
		min := l
		for i := l + 1; i < n; i++ {
			if a[i] < a[min] {
				min = i
			}
		}
		a[min], a[l] = a[l], a[min]
	}
}

func mergeSlices(a, b []int) []int {
	x, i, j := 0, 0, 0
	c := make([]int, len(a)+len(b))

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c[x] = a[i]
			i++
		} else {
			c[x] = b[j]
			j++
		}
		x++
	}
	for i < len(a) {
		c[x] = a[i]
		i++
		x++
	}
	for j < len(b) {
		c[x] = b[j]
		j++
		x++
	}
	return c
}

func MergeSort(a []int) {
	var sort func([]int) []int
	sort = func(a []int) []int {
		if len(a) < 2 {
			return a
		}

		m := len(a) / 2
		l := sort(a[m:])
		r := sort(a[:m])

		res := mergeSlices(l, r)
		return res
	}
	sorted := sort(a)
	for i := range a {
		a[i] = sorted[i]
	}
}

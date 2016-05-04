package searchsort

type HashFunction func(int, int) int

func LinearSearch(a []int, b int) bool {
	for _, x := range a {
		if x == b {
			return true
		}
	}
	return false
}

func BinarySearch(a []int, b int) bool {
	var m, l, r = 0, 0, len(a)
	for r > l {
		m = (l + r) / 2
		e := a[m]
		if b == e {
			return true
		} else if b > e {
			l = m + 1
		} else {
			r = m
		}
	}
	return false
}

func HashSearch(hash HashFunction, a []int, b int) bool {
	l := len(a)
	if l > 0 {
		h := hash(b, l)
		e := a[h]
		return e == b
	}
	return false
}

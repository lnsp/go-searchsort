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
        for i := l+1; i < n; i++ {
            if a[i] < a[min] {
                min = i
            }
        }
        a[min], a[l] = a[l], a[min]
    }
}
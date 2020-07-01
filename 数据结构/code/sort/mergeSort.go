package sort

// MergeSort 归并排序
func MergeSort(a []int) {
	if len(a) <= 1 {
		return
	}

	mergeSort(a, 0 , len(a)-1)
}

func mergeSort(a []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end ) / 2
	mergeSort(a, start, mid)
	mergeSort(a, mid+1, end)
	merge(a, start, mid, end)
}

func merge(a []int, start, mid, end int) {
	tmp := make([]int, end-start+1)
	i, j, k := start, mid+1, 0
	for i <= mid && j <= end {
		if a[i] <= a[j] {
			tmp[k]=a[i]
			i++
		} else {
			tmp[k]=a[j]
			j++
		}
		k++
	}

	for i <=mid {
		tmp[k]=a[i]
		k++
		i++
	} 

	for j <= end {
		tmp[k]=a[j]
		k++
		j++
	}
	copy(a[start:end+1], tmp)

}


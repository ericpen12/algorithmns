package sort

func MergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) >> 1
	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	i, j := left, mid+1
	tmp := make([]int, 0, right-left+1)
	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			tmp = append(tmp, arr[i])
			i++
		} else {
			tmp = append(tmp, arr[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, arr[i])
	}

	for ; j <= right; j++ {
		tmp = append(tmp, arr[j])
	}

	for i, v := range tmp {
		arr[left+i] = v
	}
}

package sort


func QuickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(arr, start, end)
	QuickSort(arr, start, pivot-1)
	QuickSort(arr, pivot+1, end)
}

func partition(arr []int, start, end int) int {
	counter, pivot := start, end
	for i := start; i <= end; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[counter] = arr[counter], arr[i]
			counter++
		}
	}
	arr[pivot], arr[counter] = arr[counter], arr[pivot]
	return counter
}
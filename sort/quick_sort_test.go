package sort

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{1,3, 2, 4, 65, 32, 43, 45, 122}
	QuickSort(arr, 0, len(arr)-1)
	for i:=1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			t.Errorf("排序失败，arr: %v", arr)
			return
		}
	}
	t.Log(arr)
}
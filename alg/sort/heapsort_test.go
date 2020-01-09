package sort

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	nums := []int{9, 4, 7, 1, 5, 3, 8, 2, 6}
	HeapSort(nums)
	if !reflect.DeepEqual(nums, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Errorf("after sort: %v", nums)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeapSort(randomSlice(b))
	}
}

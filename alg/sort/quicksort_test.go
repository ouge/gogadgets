package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{9, 4, 7, 1, 5, 3, 8, 2, 6}
	QuickSort(nums)
	if !reflect.DeepEqual(nums, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Errorf("after sort: %v", nums)
	}
}

func TestQuickSort2(t *testing.T) {
	nums := []int{9, 4, 7, 1, 5, 3, 8, 2, 6}
	QuickSort2(nums)
	if !reflect.DeepEqual(nums, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Errorf("after sort: %v", nums)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(randomSlice(b))
	}
}

func BenchmarkQuickSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort2(randomSlice(b))
	}
}

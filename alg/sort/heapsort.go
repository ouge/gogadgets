package sort

func HeapSort(values []int) {
	// 初始化堆
	for i := len(values)/2 - 1; i >= 0; i-- {
		sink(values, i)
	}
	// 排序
	for i := len(values) - 1; i >= 1; i-- {
		values[i], values[0] = values[0], values[i]
		sink(values[:i], 0)
	}
}

// 节点下沉
func sink(values []int, k int) {
	for {
		l, r := 2*k+1, 2*k+2
		if l < len(values) && r >= len(values) && values[l] > values[k] {
			// 有左孩子，无右孩子，左孩子是两者最大
			values[k], values[l] = values[l], values[k]
			k = l
		} else if r < len(values) && values[l] <= values[r] && values[k] < values[r] {
			// 有两孩子，右孩子是三者最大
			values[k], values[r] = values[r], values[k]
			k = r
		} else if r < len(values) && values[l] > values[r] && values[k] < values[l] {
			// 有两孩子，左孩子是三者最大
			values[k], values[l] = values[l], values[k]
			k = l
		} else {
			break
		}
	}
}

package sort

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}

func quickSort(values []int, left, right int) {
	if right-left < 1 {
		return
	}
	val := values[left]
	i, j := left, right
	for i < j {
		for i < j && values[j] >= val {
			j--
		}
		if i < j {
			values[i] = values[j]
		}
		for i < j && values[i] <= val {
			i++
		}
		if i < j {
			values[j] = values[i]
		}
	}
	values[i] = val
	quickSort(values, left, i-1)
	quickSort(values, i+1, right)
}

func QuickSort2(values []int) {
	if len(values) <= 1 {
		return
	}
	val := values[0]
	i, j := 0, len(values)-1
	for i < j {
		for i < j && values[j] >= val {
			j--
		}
		if i < j {
			values[i] = values[j]
		}
		for i < j && values[i] <= val {
			i++
		}
		if i < j {
			values[j] = values[i]
		}
	}
	values[i] = val
	QuickSort2(values[:i])
	QuickSort2(values[i+1:])
}

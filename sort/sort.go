package sort

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	m, j := arr[0], 1
	for i := 1; i < len(arr); i++ {
		if arr[i] <= m {
			arr[j], arr[i] = arr[i], arr[j]
			j++
		}
	}
	arr[0], arr[j-1] = arr[j-1], arr[0]

	QuickSort(arr[:j-1])
	QuickSort(arr[j:])

}

func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2
	MergeSort(arr[:mid])
	MergeSort(arr[mid:])

	brr := make([]int, 0, len(arr))
	i, j := 0, mid
	for i < mid && j < len(arr) {
		if arr[i] < arr[j] {
			brr = append(brr, arr[i])
			i++
		} else {
			brr = append(brr, arr[j])
			j++
		}
	}
	for i < mid {
		brr = append(brr, arr[i])
		i++
	}
	for j < len(arr) {
		brr = append(brr, arr[j])
		j++
	}
	for k, _ := range arr {
		arr[k] = brr[k]
	}
}

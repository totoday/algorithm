package sort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{3, 1, 2}
	MergeSort(arr)
	fmt.Println(arr)
}
package algorithm

import "testing"

func TestInsertionSort(t *testing.T) {
	in := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := InsertionSort(in)
	for i, v := range want {
		if v != got[i] {
			t.Errorf("InsertionSort(%v) = %v; want %v", in, got, want)
			break
		}
	}
}

func TestMergeSort(t *testing.T) {
	in := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := MergeSort(in)
	for i, v := range want {
		if v != got[i] {
			t.Errorf("MergeSort(%v) = %v; want %v", in, got, want)
			break
		}
	}
}

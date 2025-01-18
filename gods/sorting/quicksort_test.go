package sorting

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	// t.Fatal("not implemented")

	arr := []int{9, 12, 1, 15, 6}
	want := []int{1, 6, 9, 12, 15}
	quicksort(&arr, 0, len(arr)-1)
	if !reflect.DeepEqual(arr, want) {
		t.Errorf("got %v   want %v", arr, want)
	}
}

func TestQuickSortWithOneElemet(t *testing.T) {
	// t.Fatal("not implemented")

	arr := []int{9}
	want := []int{9}
	quicksort(&arr, 0, len(arr)-1)
	if !reflect.DeepEqual(arr, want) {
		t.Errorf("got %v   want %v", arr, want)
	}
}

func TestQuickSortWithTwoElemet(t *testing.T) {
	// t.Fatal("not implemented")

	arr := []int{9, 1}
	want := []int{1, 9}
	quicksort(&arr, 0, len(arr)-1)
	if !reflect.DeepEqual(arr, want) {
		t.Errorf("got %v   want %v", arr, want)
	}
}

func TestQuickSortWithSortedArray(t *testing.T) {
	// t.Fatal("not implemented")

	arr := []int{1, 2, 14, 20}
	want := []int{1, 2, 14, 20}
	quicksort(&arr, 0, len(arr)-1)
	if !reflect.DeepEqual(arr, want) {
		t.Errorf("got %v   want %v", arr, want)
	}
}

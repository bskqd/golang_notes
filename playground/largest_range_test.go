package main

import (
	"reflect"
	"testing"
)

func TestLargestRange(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 0, 3, 4, 5}
	largestRangeArr1 := largestRange(arr1)
	largestRangeArr2 := largestRange(arr2)
	if !reflect.DeepEqual(largestRangeArr1, arr1) {
		t.Fatalf("largestRangeArr1: %v, expected return: %v", largestRangeArr1, arr1)
	}
	if !reflect.DeepEqual(largestRangeArr2, []int{3, 4, 5}) {
		t.Fatalf("largestRangeArr2: %v, expected return: %v", largestRangeArr2, []int{3, 4, 5})
	}
}

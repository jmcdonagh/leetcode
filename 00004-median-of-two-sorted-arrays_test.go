package main

import (
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"[1,3] [2]", args{ []int{1,3}, []int{2} }, 2.00000 },
		{"[1,3,5,11,11,12] [2,4,6,7,8,9,10]", args{ []int{1,3,5,11,11,12}, []int{2,4,6,7,8,9,10} }, 7.00000 },
		{"[1,2] [3,4]", args{ []int{1,2}, []int{3,4} }, 2.50000 },
		{"[0,0,0,0,0] [-1,0,0,0,0,0,1]", args{ []int{0,0,0,0,0}, []int{-1,0,0,0,0,0,1} }, 0 },
		{"[1,2] [1,2]", args{ []int{1,2}, []int{1,2} }, 1.50000 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

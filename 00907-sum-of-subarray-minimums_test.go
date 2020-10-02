package main

import (
	"reflect"
	"testing"
)

func Test_sumSubarrayMins(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{ "[3,1,2,4]", args{[]int{3,1,2,4}}, 17 },
		{"[1,2,3,4,5]", args{[]int{1,2,3,4,5}}, 35 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSubarrayMins(tt.args.A); got != tt.want {
				t.Errorf("sumSubarrayMins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSubArrays(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"{3,1,2,4}", args{[]int{3,1,2,4}}, [][]int{{3},{1},{2},{4},{3,1},{1,2},{2,4},{3,1,2},{1,2,4},{3,1,2,4}} },
		{"{1,2,3,4,5}", args{[]int{1,2,3,4,5}}, [][]int{{1},{2},{3},{4},{5},{1,2},{2,3},{3,4},{4,5},{1,2,3},{2,3,4},{3,4,5},{1,2,3,4},{2,3,4,5},{1,2,3,4,5}} },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSubArrays(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSubArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

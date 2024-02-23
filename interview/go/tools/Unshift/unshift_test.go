package main

import (
	"reflect"
	"testing"
)

func TestPreAppend(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		values []int
		want   []int
	}{
		{
			name:   "prepend 1, 2, 3 to empty slices",
			nums:   []int{},
			values: []int{1, 2, 3},
			want:   []int{3, 2, 1},
		},
		{
			name:   "prepend 1, 2, 3 to existing slices",
			nums:   []int{5, 4},
			values: []int{1, 2, 3},
			want:   []int{3, 2, 1, 5, 4},
		},
	}

	for _, tt := range testCases {
		if got := prependNums(tt.nums, tt.values...); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("prepend result is %v, but got %v", got, tt.want)
		}
	}
}

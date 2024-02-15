package main

import (
	"reflect"
	"testing"
)

func TestPreAppend(t *testing.T) {
	test := []struct {
		name   string
		nums   []int
		values []int
		want   []int
	}{
		{
			name:   "prepend 1, 2, 3 to empty slice",
			nums:   []int{},
			values: []int{1, 2, 3},
			want:   []int{3, 2, 1},
		},
		{
			name:   "prepend 1, 2, 3 to existing slice",
			nums:   []int{5, 5},
			values: []int{1, 2, 3},
			want:   []int{3, 2, 1, 5, 5},
		},
	}

	for _, tt := range test {
		if got := prependNums(tt.nums, tt.values...); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("prepend result is %v, but want %v", got, tt.want)
		}
	}
}

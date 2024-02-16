package main

import (
	"reflect"
	"testing"
)

func TestLIS(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test1",
			input: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want:  4,
		},
		{
			name:  "test2",
			input: []int{0, 1, 0, 3, 2, 3},
			want:  4,
		},
		{
			name:  "test3",
			input: []int{7, 7, 7, 7, 7},
			want:  1,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.input); !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("%v 's answer is %v, but got %v", tt.name, tt.want, got)
			}
		})
	}
}

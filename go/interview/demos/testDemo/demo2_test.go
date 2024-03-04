package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name   string
		input1 int
		input2 int
		want   int
	}{
		{
			name:   "test case 1",
			input1: 8,
			input2: 15,
			want:   23,
		},
		{
			name:   "test case 2",
			input1: -5,
			input2: 10,
			want:   5,
		},
	}

	for _, tt := range testCases {
		if got := Add(tt.input1, tt.input2); got != tt.want {
			t.Errorf("%v:\n prepend Result is %v, but got %v", tt.name, tt.want, got)
		}
	}
}

func TestJoin(t *testing.T) {
	testCases := []struct {
		name   string
		input1 int
		input2 int
		want   []int
	}{
		{
			name:   "test case 1",
			input1: 2,
			input2: 0,
			want:   []int{0, 0},
		},
		{
			name:   "test case 2",
			input1: 3,
			input2: -1,
			want:   []int{-1, -1, -1},
		},
	}

	for _, tc := range testCases {
		if got := Join(tc.input1, tc.input2); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("\n %v: \n  want: %v, but got %v", tc.name, tc.want, got)
		}
	}
}

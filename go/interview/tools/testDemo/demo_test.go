package main

import "testing"

func TestAdd(t *testing.T) {
	ans := Add(1, 2)

	if ans != 3 {
		t.Errorf("1 + 2 excepted be 3! But your answer is: %d", ans)
	}

	ans = Add(-10, -20)
	if ans != -30 {
		t.Errorf("-10 + -20 excepted be -30! But your answer is: %d", ans)
	}
}

func TestMul(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if Mul(2, 3) != 6 {
			t.Errorf("fail")
		}
	})

	t.Run("neg", func(t *testing.T) {
		if Mul(-2, -10) != 20 {
			t.Errorf("fail")
		}
	})

}

func TestDemo(t *testing.T) {
	testCaess := []struct {
		name string
		num1 int
		num2 int
		want int
	}{
		{
			name: "multiple: 5 * 5",
			num1: 5,
			num2: 5,
			want: 25,
		},
		{
			name: "multiple: 15 * 7",
			num1: 15,
			num2: 7,
			want: 105,
		},
	}

	for _, tt := range testCaess {
		if got := Mul(tt.num1, tt.num2); got != tt.want {
			t.Errorf("prepend Result is %v, but got %v", tt.want, got)
		}
	}
}

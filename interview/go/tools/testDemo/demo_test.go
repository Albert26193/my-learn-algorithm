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
			t.Fatal("fail")
		}
	})

	t.Run("neg", func(t *testing.T) {
		if Mul(-2, -10) != 20 {
			t.Fatal("fail")
		}
	})
}

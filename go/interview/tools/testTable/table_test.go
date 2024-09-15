package main

import "testing"

func TestMul(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Excepted int
	}{
		{"pos", 2, 3, 6},
		{"neg", 2, -3, -6},
		{"zero", 2, 0, 0},
	}
			
	for _, c := range cases {
		if ans := Mul(c.A, c.B); ans != c.Excepted {
			t.Fatalf("%d * %d should be %d", c.A, c.B, c.Excepted)
		}
	}
}

func TestAgain(t *testing.T) {
	cases := []struct {
		Name string
		Num1, Num2, Wish int
	} {
		{"test1", 1, 2, 3},
		{"test2", -1, 12, 11},
		{"test3", 2, 0, 2},
	}

	for _, c := range cases {
		if ans := Add(c.Num1, c.Num2); ans != c.Wish {
			t.Errorf("%s, %d + %d should be %d, but got %d", c.Name, c.Num1, c.Num2, c.Wish, ans)
		}
	}
}
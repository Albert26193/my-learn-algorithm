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
		t.Run(c.Name, func(t *testing.T) {
			if ans := Mul(c.A, c.B); ans != c.Excepted {
				t.Fatalf("%d * %d should be %d", c.A, c.B, c.Excepted)
			}
		})
	}
}

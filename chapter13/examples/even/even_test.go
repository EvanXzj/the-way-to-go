package even

import "testing"

func TestEven(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{10, true},
		{20, true},
		{21, false},
	}

	for _, c := range cases {
		got := Even(c.in)
		if got != c.want {
			t.Errorf("Even(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestOdd(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{10, false},
		{20, false},
		{21, true},
	}
	for _, c := range cases {
		got := Odd(c.in)
		if got != c.want {
			t.Errorf("Odd(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

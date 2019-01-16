package even

import "testing"

func TestEven(t *testing.T) {
	cases := []struct {
		in int
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
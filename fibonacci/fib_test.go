package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	var tests = []struct {
		name string
		n    uint
		want uint64
	}{
		{"1", 1, 1},
		{"2", 2, 1},
		{"3", 3, 2},
		{"17", 17, 1597},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got, _ := Fibonacci(test.n); got != test.want {
				t.Errorf("Fibonacci(%v), got %v, want %v", test.n, got, test.want)
			}
		})
	}

}

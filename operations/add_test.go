package operations

import "testing"

func TestAdd(t *testing.T) {
	i, j := 2, 3
	want := 5
	got := Add(i, j)
	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

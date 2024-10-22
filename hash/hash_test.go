package hash

import "testing"

func TestCanConstruct(t *testing.T) {
	got := CanConstruct("aa", "ab")
	want := false

	if got != want {
		t.Errorf("Uncorect algorythm")
	}
}

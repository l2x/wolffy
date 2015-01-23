package utils

import "testing"

func TestDelEmptySlice(t *testing.T) {
	s := []string{"", "a", "c", "b"}
	s = DelEmptySlice(s)

	if len(s) != 3 {
		t.Error("length error")
	}

	for _, v := range s {
		if v == "" {
			t.Error("delete empty error")
		}
	}
}

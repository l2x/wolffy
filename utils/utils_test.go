package utils

import (
	"sort"
	"testing"
)

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

func TestStringReverse(t *testing.T) {
	s := []string{"c", "a", "b"}
	ss := []string{"c", "b", "a"}

	sort.Sort(StringReverse(s))

	if !compareStringSlice(s, ss) {
		t.Errorf("reverse error %v %v", s, ss)
	}
}

func compareStringSlice(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for k, v := range s1 {
		if v != s2[k] {
			return false
		}
	}

	return true
}

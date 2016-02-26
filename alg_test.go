package alg

import (
	"testing"
)

func TestLevenshtein(t *testing.T) {
	a := "FOOD"
	b := "MONEY"
	if l := Levenshtein(a, b); l != 4 {
		t.Errorf("Expected 4, got %d", l)
	}
}

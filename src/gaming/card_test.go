package gaming

import (
	"testing"
)

func TestSuits(t *testing.T) {
	if len(Suits()) != 4 {
		t.Error("Expect 4 suits")
	}
}

func TestValues(t *testing.T) {
	if len(Values()) != 13 {
		t.Error("Expect 13 values")
	}
	if len(JokerValues()) != 14 {
		t.Error("Expect 14 joker values")
	}
}

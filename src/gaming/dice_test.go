package gaming

import (
	"testing"
)

func TestRoll(t *testing.T) {
	d1 := NewDiceSeed(uint(6), 10)
	d2 := NewDiceSeed(uint(6), 10)
	if d1.Roll() != d2.Roll() {
		t.Error("Expected dice to roll the same with the same seed")
	}
	if d1.Sides() != d2.Sides() {
		t.Error("Expected dice to have the same number of sides")
	}
}

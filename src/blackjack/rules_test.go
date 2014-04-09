package blackjack

import (
	"gaming"
	"testing"
)

func TestRules(t *testing.T) {
	rules := NewRulesetFactory().HardDoubleValues([]uint{9, 10, 11}).Build()
	h := NewHand()
	h.Push(NewCard(gaming.Spade, Four))
	h.Push(NewCard(gaming.Spade, Four))

	if rules.CanDouble(h) {
		t.Error("Cannot double on 8!")
	}
	h = NewHand()
	h.Push(NewCard(gaming.Spade, Four))
	h.Push(NewCard(gaming.Spade, Five))
	if !rules.CanDouble(h) {
		t.Error("Can double on 9!")
	}
	if rules.BlackjackPayout() != NORMAL_PAYOUT {
		t.Error("Expected normal blackjack payout")
	}
}

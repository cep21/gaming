package blackjack

import (
	"testing"
	"gaming"
)

func TestRules(t *testing.T) {
	rules := NewRuleset(true, true, true, []uint{9, 10, 11}, []uint{9, 10, 11}, NORMAL_PAYOUT, .5)
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

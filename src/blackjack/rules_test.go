package blackjack

import (
	"testing"
)

func TestRules(t *testing.T) {
	rules := NewRuleset(true, true, true, []uint{9, 10, 11}, NORMAL_PAYOUT)
	if rules.CanDouble(uint(8)) {
		t.Error("Cannot double on 8!")
	}
	if !rules.CanDouble(uint(9)) {
		t.Error("Can double on 9!")
	}
}

package blackjack

import "testing"

func TestDeck(t *testing.T) {
	if SingleDeck().CardsLeft() != 52 {
		t.Errorf("Expect 52 values")
	}
}

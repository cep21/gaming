package gaming

import "testing"

func TestDeck(t *testing.T) {
	if SingleDeck().CardsLeft() != 52 {
		t.Errorf("Expect 52 values")
	}
	if JokerDeck(1).CardsLeft() != 53 {
		t.Errorf("Expect 53 joker values")
	}
}

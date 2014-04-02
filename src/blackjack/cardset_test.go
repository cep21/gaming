package blackjack

import (
	"testing"
	"math/rand"
	"gaming"
)

func TestDeck(t *testing.T) {
	if SingleDeck().CardsLeft() != 52 {
		t.Errorf("Expect 52 values")
	}
}

func TestInfiniteDeck(t *testing.T) {
	r := rand.New(rand.NewSource(2))
	d := NewInfiniteDeck(r)
	if d.IsEmpty() {
		t.Error("Infinite decks are never empty")
	}
	if d.CardsLeft() < 10000 {
		t.Error("Deck should be infinite")
	}
	// Consistent because seed is a constant
	if d.Pop().Value() == d.Pop().Value() {
		t.Error("Expected different values when I pop'd")
	}
	for i := 0; i < 1000; i++ {
		d.Pop()
	}
	c := d.Pop()
	d.Push(c)
	if c.Value() != d.Pop().Value() {
		t.Error("Pushing back on an infinite deck does not work")
	}
}

func TestRandomPickDeck(t *testing.T) {
	r := rand.New(rand.NewSource(2))
	d := NewRandomPickDeck(r, gaming.Spade, 1)
	cards_left := uint(52);
	num_eights := 0;
	for i:=0;i<52;i++ {
		c := d.Pop()
		if c.Value() == Eight {
			num_eights++
		}
		cards_left--;
		if d.CardsLeft() != cards_left {
			t.Errorf("Expected %d cards left, not %d\n", cards_left, d.CardsLeft())
		}
	}
	if num_eights != 4 {
		t.Error("Expected 4 eights")
	}

}

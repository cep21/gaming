/**
 * Date: 4/3/14
 * Time: 1:24 PM
 * @author jack 
 */
package blackjack

import (
	"testing"
	"math/rand"
)

func TestDeck(t *testing.T) {
	d := Decks(1)
	singleDeckVerificiation(t, d.Clone(), 1)
	singleDeckVerificiation(t, d, 1)
}

func TestInfiniteDeck(t *testing.T) {
	r := rand.New(rand.NewSource(2))
	d := NewInfiniteShoe(r)
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
	if c == nil {
		t.Error("Do not expect nil cards from an infinite shoe")
	}
	for i := 0;i<1000;i++ {
		if d.TakeValueFromShoe(Ace) == nil {
			t.Error("Did not expect to run out of aces from an infinite shoe")
		}
	}
}

func TestRandomPickDeck(t *testing.T) {
	r := rand.New(rand.NewSource(2))
	d := NewRandomPickShoe(r, 1)
	singleDeckVerificiation(t, d.Clone().Shuffle(r), 1)
	singleDeckVerificiation(t, d, 1)
}

func singleDeckVerificiation(t *testing.T, d Shoe, number_of_decks uint) {
	if d.CardsLeft() != 52 {
		t.Errorf("Expect 52 values")
	}
	counts_per_suit := []uint{0, 0, 0, 0}
	for i := 0;i<4;i++ {
		c := d.TakeValueFromShoe(Ace)
		if c == nil {
			t.Error("Did not expect to run out of 4 aces")
		} else if c.Value() != Ace {
			t.Error("Did not draw an ace")
		}
		counts_per_suit[c.Suit().Index()]++
	}
	for _, v := range counts_per_suit {
		if v != number_of_decks {
			t.Error("Did not draw the right amount of a suit from a deck")
		}
	}

	if d.TakeValueFromShoe(Ace) != nil {
		t.Error("We should be out of aces by now!")
	}

	cards_left := uint(52 - 4);
	num_eights := 0;
	for i:=0;i<52-4;i++ {
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

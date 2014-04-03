/**
 * Date: 4/3/14
 * Time: 1:24 PM
 * @author jack 
 */
package blackjack

import (
	"testing"
	"math/rand"
	"gaming"
)

func TestDeck(t *testing.T) {
	if Decks(1).CardsLeft() != 52 {
		t.Errorf("Expect 52 values")
	}
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
}

func TestRandomPickDeck(t *testing.T) {
	r := rand.New(rand.NewSource(2))
	d := NewRandomPickShoe(r, gaming.Spade, 1)
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

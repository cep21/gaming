/**
 * Date: 4/4/14
 * Time: 1:24 PM
 * @author jack 
 */
package blackjack

import (
	"testing"
	"gaming"
)

func TestBasicHandDealer(t *testing.T) {
	shoe := Decks(1)
	dealer := NewBasicHandDealer()
	hands, _ := dealer.DealHands(shoe, 2)
	basicHandVerification(t, hands, 2)
}

func TestForcedHandsDealer(t *testing.T) {
	shoe := Decks(1)
	dealerUpCard := Ten
	playerHand := NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Ten))
	dealer := NewForceDealerPlayerHands(playerHand, dealerUpCard)
	hands, _ := dealer.DealHands(shoe, 2)
	basicHandVerification(t, hands, 2)

	// Shouldn't be able to deal ten/ten/ten again on a one deck shoe
	_, err := dealer.DealHands(shoe, 2)
	if err == nil {
		t.Error("I expected an error dealing this hand twice")
	}

}

func basicHandVerification(t *testing.T, hands []Hand, expectedSize uint) {
	if uint(len(hands)) != expectedSize {
		t.Error("Unexpected number of hands")
	}
	for _, h := range hands {
		if h.Size() != uint(2) {
			t.Error("Dealt hands should have two cards")
		}
	}
}

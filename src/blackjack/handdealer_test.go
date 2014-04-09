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
	shoe := Decks(uint(1))
	dealer := NewDealer(NewDealerStrategy(false), NewBasicHandDealer())
	player := NewPlayer(NewConsistentBettingStrategy(1), NewAlwaysHitStrategy())
	players := []Player{player}
	err := dealer.HandDealer().DealHands(shoe, players, dealer)
	if err != nil {
		t.Fatal("Don't expect failed deals")
	}
	basicHandVerification(t, players, dealer)
}

func TestForcedHandsDealer(t *testing.T) {

	shoe := Decks(uint(1))
	dealerUpCard := Ten
	playerHand := NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Ten))
	dealer := NewDealer(NewDealerStrategy(false), NewForceDealerPlayerHands(playerHand, dealerUpCard))
	player := NewPlayer(NewConsistentBettingStrategy(1), NewAlwaysHitStrategy())
	players := []Player{player}
	err := dealer.HandDealer().DealHands(shoe, players, dealer)
	if err != nil {
		t.Fatal("Don't expect failed deals")
	}
	basicHandVerification(t, players, dealer)

	// Shouldn't be able to deal ten/ten/ten again on a one deck shoe
	err = dealer.HandDealer().DealHands(shoe, players, dealer)
	if err == nil {
		t.Error("I expected an error dealing this hand twice")
	}
}

func basicHandVerification(t *testing.T, players []Player, dealer Dealer) {
	for _, p := range players {
		for _, h := range p.Hands() {
			if h.Size() != uint(2) {
				t.Error("Dealt hands should have two cards")
			}
		}
	}
	for _, h := range dealer.Hands() {
		if h.Size() != uint(2) {
			t.Error("Dealt hands should have two cards")
		}
	}
}

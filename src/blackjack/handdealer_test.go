/**
 * Date: 4/4/14
 * Time: 1:24 PM
 * @author jack 
 */
package blackjack

import (
	"testing"
	"gaming"
	"gaming/bankroll"
)

func TestBasicHandDealer(t *testing.T) {
	shoe := Decks((uint)(1))
	strat := NewConsistentBettingStrategy(1)
	bank := bankroll.NewMoneyHolder()
	dealer := NewBasicHandDealer()
	player_hands, dealer_hand, err := dealer.DealHands(shoe, []BettingStrategy{strat}, []bankroll.MoneyHolder{bank})
	if err != nil {
		t.Fatal("Don't expect failed deals")
	}
	basicHandVerification(t, player_hands, dealer_hand, (uint)(1))
}

func TestForcedHandsDealer(t *testing.T) {
	shoe := Decks((uint)(1))
	dealerUpCard := Ten
	playerHand := NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Ten))
	strat := NewConsistentBettingStrategy(1)
	bank := bankroll.NewMoneyHolder()
	dealer := NewForceDealerPlayerHands(playerHand, dealerUpCard)
	player_hands, dealer_hand, err := dealer.DealHands(shoe, []BettingStrategy{strat}, []bankroll.MoneyHolder{bank})
	if err != nil {
		t.Fatal("Don't expect failed deals")
	}
	basicHandVerification(t, player_hands, dealer_hand, (uint)(1))

	// Shouldn't be able to deal ten/ten/ten again on a one deck shoe
	_, _, err = dealer.DealHands(shoe, []BettingStrategy{strat}, []bankroll.MoneyHolder{bank})
	if err == nil {
		t.Error("I expected an error dealing this hand twice")
	}

}

func basicHandVerification(t *testing.T, player_hands []Hand, dealer_hand Hand, expectedSize uint) {
	if uint(len(player_hands)) != expectedSize {
		t.Error("Unexpected number of hands")
	}
	player_hands = append(player_hands, dealer_hand)
	for _, h := range player_hands {
		if h.Size() != uint(2) {
			t.Error("Dealt hands should have two cards")
		}
	}
}

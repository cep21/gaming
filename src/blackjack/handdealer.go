/**
 * Date: 4/4/14
 * Time: 1:10 PM
 * @author jack 
 */
package blackjack

import (
	"errors"
)

var ERR_NOT_ENOUGH_CARDS_TO_DEAL_HANDS = errors.New("Not enough cards to deal all the hands")

type HandDealer interface {
	DealHands(deck Shoe, players []Player, dealer HandHolder) error
}

type basicHandDealer struct {
}

func NewBasicHandDealer() HandDealer {
	return &basicHandDealer{}
}

func (this *basicHandDealer) DealHands(deck Shoe, players []Player, dealer HandHolder) error {
	player_hands := make([]Hand, len(players))
	dealer_hand := NewHand()
	dealer.SetHands([]Hand{dealer_hand})
	for j := 0; j < 2; j++ {
		for i := 0; i < len(players); i++ {
			if j == 0 {
				player_hands[i] = NewHand()
				players[i].SetHands([]Hand{player_hands[i]})
				units_to_bet := players[i].BettingStrategy().GetMoneyToBet()
				players[i].Bankroll().TransferMoneyTo(player_hands[i].MoneyInThisHand(), units_to_bet)
			}
			card, err := deck.Pop()
			if err != nil {
				return err
			}
			player_hands[i].Push(card)
		}
		card, err := deck.Pop()
		if err != nil {
			return err
		}
		dealer_hand.Push(card)
	}
	return nil
}

type forceDealerPlayerHands struct {
	playerHandToForce   Hand
	dealerUpCardToForce Value
}

func NewForceDealerPlayerHands(playerHandToForce Hand, dealerUpCardToForce Value) HandDealer {
	return &forceDealerPlayerHands{
		playerHandToForce: playerHandToForce,
		dealerUpCardToForce: dealerUpCardToForce,
	}
}

func (this *forceDealerPlayerHands) DealHands(deck Shoe, players []Player, dealer HandHolder) error {
	for i := 0; i < len(players); i++ {
		hand_cards := []Card{}
		for _, c := range this.playerHandToForce.Cards() {
			card, err := deck.TakeValueFromShoe(c.BlackjackValue())
			if err != nil {
				return err
			}
			hand_cards = append(hand_cards, card)
		}

		playerHand := NewHand(hand_cards...)
		units_to_bet := players[i].BettingStrategy().GetMoneyToBet()
		players[i].Bankroll().TransferMoneyTo(playerHand.MoneyInThisHand(), units_to_bet)
		players[i].SetHands([]Hand{playerHand})
	}
	dealerFirstCard, err := deck.TakeValueFromShoe(this.dealerUpCardToForce)
	if err != nil {
		return err
	}
	dealerHand := NewHand(dealerFirstCard)
	dealer.SetHands([]Hand{dealerHand})
	dealer_second_card, err := deck.Pop()
	if err != nil {
		return err
	}
	dealerHand.Push(dealer_second_card)
	return nil
}


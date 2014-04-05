/**
 * Date: 4/4/14
 * Time: 1:10 PM
 * @author jack 
 */
package blackjack

import "errors"

var ERR_NOT_ENOUGH_CARDS_TO_DEAL_HANDS = errors.New("Not enough cards to deal all the hands")

type HandDealer interface {
	DealHands(deck Shoe, number_of_players uint) ([]Hand, Hand, error)
}

type basicHandDealer struct {
}

func NewBasicHandDealer() HandDealer {
	return &basicHandDealer{}
}


func (this *basicHandDealer) DealHands(deck Shoe, number_of_players uint) ([]Hand, Hand, error) {
	player_hands := make([]Hand, number_of_players)
	dealer_hand := NewHand()
	for j:=0;j<2;j++ {
		for i := uint(0); i<number_of_players;i++ {
			if j == 0 {
				player_hands[i] = NewHand()
			}
			card, err := deck.Pop()
			if err != nil {
				return nil, nil, err
			}
			player_hands[i].Push(card)
		}
		card, err := deck.Pop()
		if err != nil {
			return nil, nil, err
		}
		dealer_hand.Push(card)
	}
	return player_hands, dealer_hand, nil
}

type forceDealerPlayerHands struct {
	playerHandToForce Hand
	dealerUpCardToForce Value
}

func NewForceDealerPlayerHands(playerHandToForce Hand, dealerUpCardToForce Value) HandDealer {
	return &forceDealerPlayerHands{
		playerHandToForce: playerHandToForce,
		dealerUpCardToForce: dealerUpCardToForce,
	}
}

func (this *forceDealerPlayerHands) DealHands(deck Shoe, number_of_players uint) ([]Hand, Hand, error) {
	player_hands := make([]Hand, number_of_players)
	dealer_hand := NewHand()

	for i := uint(0); i<number_of_players;i++ {
		hand_cards := []Card{}
		for _, c := range this.playerHandToForce.Cards() {
			card, err := deck.TakeValueFromShoe(c.BlackjackValue())
			if err != nil {
				return nil, nil, err
			}
			hand_cards = append(hand_cards, card)
		}
		player_hands[i] = NewHand(hand_cards...)
	}
	card, err := deck.TakeValueFromShoe(this.dealerUpCardToForce)
	if err != nil {
		return nil, nil, err
	}
	dealer_hand = NewHand(card)
	dealer_second_card, err := deck.Pop()
	if err != nil {
		return nil, nil, err
	}
	dealer_hand.Push(dealer_second_card)
	return player_hands, dealer_hand, nil
}

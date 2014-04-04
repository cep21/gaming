/**
 * Date: 4/4/14
 * Time: 1:10 PM
 * @author jack 
 */
package blackjack

import "errors"

var ERR_NOT_ENOUGH_CARDS_TO_DEAL_HANDS = errors.New("Not enough cards to deal all the hands")

type HandDealer interface {
	DealHands(deck Shoe, number_of_hands uint) ([]Hand, error)
}

type basicHandDealer struct {
}

func NewBasicHandDealer() HandDealer {
	return &basicHandDealer{}
}


func (this *basicHandDealer) DealHands(deck Shoe, number_of_hands uint) ([]Hand, error) {
	hands := make([]Hand, number_of_hands)
	for j:=0;j<2;j++ {
		for i := uint(0); i<number_of_hands;i++ {
			if j == 0 {
				hands[i] = NewHand()
			}
			card, err := deck.Pop()
			if err != nil {
				return nil, err
			}
			hands[i].Push(card)
		}
	}
	return hands, nil
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

func (this *forceDealerPlayerHands) DealHands(deck Shoe, number_of_hands uint) ([]Hand, error) {
	hands := make([]Hand, number_of_hands)
	for i := uint(0); i<number_of_hands;i++ {
		if i == number_of_hands - 1 {
			dealer_card, err := deck.TakeValueFromShoe(this.dealerUpCardToForce)
			if err != nil {
				return nil, err
			}
			dealer_second_card, err := deck.TakeValueFromShoe(this.dealerUpCardToForce)
			if err != nil {
				return nil, err
			}
			hands[i] = NewHand(dealer_card, dealer_second_card)
		} else {
			hand_cards := []Card{}
			for _, c := range this.playerHandToForce.Cards() {
				card, err := deck.TakeValueFromShoe(c.BlackjackValue())
				if err != nil {
					return nil, err
				}
				hand_cards = append(hand_cards, card)
			}
			hands[i] = NewHand(hand_cards...)
		}
	}
	return hands, nil
}


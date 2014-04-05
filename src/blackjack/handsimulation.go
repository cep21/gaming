/**
 * Date: 4/2/14
 * Time: 2:02 PM
 * @author jack 
 */
package blackjack

import (
	"errors"
)

var ERR_TRIED_TO_SURRENDER_BUT_NOT_ALLOWED = errors.New("Tried to surrender but surrender is not allowed")
var ERR_TRIED_TO_DOUBLE_BUT_NOT_ALLOWED = errors.New("Tried to double but double is not allowed")
var ERR_TRIED_TO_SPLIT_BUT_NOT_ALLOWED = errors.New("Tried to split but split is not allowed")

func SimulateSingleHand2(shoeFactory ShoeFactory, dealer HandDealer, dealerStrategy PlayStrategy, playerStrategy PlayStrategy, bettingStrategy BettingStrategy, number_of_iterations uint, rules Rules) (float64, error) {
	bankroll := NewBankroll(0)
	for i := uint(0); i < number_of_iterations; i++ {
		units_to_bet := bettingStrategy.GetUnitsToBet()
		shoe := shoeFactory.CreateShoe()
		bankroll.ChangeBankroll(-units_to_bet)
		player_hands, dealer_hand, err := dealer.DealHands(shoe, 1)
		if err != nil {
			return 0, err
		}
		player_hand := player_hands[0]
		if dealer_hand.IsBlackjack() {
			if player_hand.IsBlackjack() {
				// Push
				bankroll.ChangeBankroll(units_to_bet)
				continue
			} else {
				continue
			}
		} else if player_hand.IsBlackjack() {
			bankroll.ChangeBankroll(+units_to_bet)
			bankroll.ChangeBankroll(units_to_bet*float64(rules.BlackjackPayout()))
			continue
		}
		hand_is_double := false
		for ; ; {
			action := playerStrategy.TakeAction(player_hand, dealer_hand.FirstCard())
			if action == SURRENDER {
				if !rules.CanSurrender(player_hand) {
					return 0, ERR_TRIED_TO_SURRENDER_BUT_NOT_ALLOWED
				}
				bankroll.ChangeBankroll(units_to_bet/2)
				break
			} else if action == DOUBLE {
				if !rules.CanDouble(player_hand) {
					return 0, ERR_TRIED_TO_DOUBLE_BUT_NOT_ALLOWED
				}
				bankroll.ChangeBankroll(-units_to_bet)
				hand_is_double = true
				card, err := shoe.Pop()
				if err != nil {
					return 0, err
				}
				player_hand.Push(card)
			} else if action == SPLIT {
				if !rules.CanSplit(player_hand) {
					return 0, ERR_TRIED_TO_SPLIT_BUT_NOT_ALLOWED
				}
				panic("I don't have split logic yet")
			} else if action == STAND {
				break
				// Do nothing.  Go on to dealer
			} else if action == HIT {
				card, err := shoe.Pop()
				if err != nil {
					return 0, err
				}
				player_hand.Push(card)
			} else {
				panic("I don't know this other thing to do ...")
			}
			if player_hand.Bust() {
				break
			}
		}
		if player_hand.Bust() {
			continue
		}
		for ; ; {
			action := dealerStrategy.TakeAction(dealer_hand, nil)
			if action == HIT {
				card, err := shoe.Pop()
				if err != nil {
					return 0, err
				}
				dealer_hand.Push(card)
				if dealer_hand.Bust() {
					break
				}
			} else if action == STAND {
				break
			} else {
				panic("Unknown dealer action!")
			}
		}
		if dealer_hand.Bust() {
			bankroll.ChangeBankroll(units_to_bet*2)
			continue
		}
		//log.Printf("%s vs %s\n", player_hand, dealer_hand)
		if dealer_hand.Score() > player_hand.Score() {
			// You loose: nothing
		} else if dealer_hand.Score() < player_hand.Score() {
			if hand_is_double {
				bankroll.ChangeBankroll(units_to_bet*4)
			} else {
				bankroll.ChangeBankroll(units_to_bet*2)
			}

		} else {
			if hand_is_double {
				bankroll.ChangeBankroll(units_to_bet*2)
			} else {
				bankroll.ChangeBankroll(units_to_bet)
			}
		}

	}
	return bankroll.CurrentBankroll()/float64(number_of_iterations), nil
}

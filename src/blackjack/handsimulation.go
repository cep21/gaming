/**
 * Date: 4/2/14
 * Time: 2:02 PM
 * @author jack
 */
package blackjack

import (
	"errors"
	"gaming/bankroll"
)

var ERR_TRIED_TO_SURRENDER_BUT_NOT_ALLOWED = errors.New("Tried to surrender but surrender is not allowed")
var ERR_TRIED_TO_DOUBLE_BUT_NOT_ALLOWED = errors.New("Tried to double but double is not allowed")
var ERR_TRIED_TO_SPLIT_BUT_NOT_ALLOWED = errors.New("Tried to split but split is not allowed")
var ERR_INVALID_DEALER_OPERATION = errors.New("Invalid dealer operations")

func PlayHand(playerHand Hand, dealerHand Hand, bankrolledStrategy BankrolledStrategy, rules Rules, houseBankroll bankroll.MoneyHolder, shoe Shoe) ([]Hand, error) {

	if len(playerHand.Cards()) == 1 {
		if playerHand.SplitNumber() == 0 {
			panic("Logic error: Should be a split hand if cards == 1")
		}
		// Force the first card
		card, err := shoe.Pop()
		if err != nil {
			return nil, err
		}
		playerHand.Push(card)
		// TODO: Verify resplits/rehits if aces/etc
	}
	for {
		var action GameAction
		if dealerHand != nil {
			action = bankrolledStrategy.PlayStrategy().TakeAction(playerHand, dealerHand.FirstCard())
		} else {
			action = bankrolledStrategy.PlayStrategy().TakeAction(playerHand, nil)
		}

		playerHand.SetLastAction(action)
		if action == SURRENDER {
			if !rules.CanSurrender(playerHand) {
				return nil, ERR_TRIED_TO_SURRENDER_BUT_NOT_ALLOWED
			}
			playerHand.MoneyInThisHand().TransferMoneyTo(bankrolledStrategy.Bankroll(), playerHand.MoneyInThisHand().CurrentBankroll()/2)
			playerHand.MoneyInThisHand().TransferMoneyTo(houseBankroll, playerHand.MoneyInThisHand().CurrentBankroll())
			break
		} else if action == DOUBLE {
			if !rules.CanDouble(playerHand) {
				return nil, ERR_TRIED_TO_DOUBLE_BUT_NOT_ALLOWED
			}
			bankrolledStrategy.Bankroll().TransferMoneyTo(playerHand.MoneyInThisHand(), playerHand.MoneyInThisHand().CurrentBankroll())
			card, err := shoe.Pop()
			if err != nil {
				return nil, err
			}
			playerHand.Push(card)
			break
		} else if action == SPLIT {
			if !rules.CanSplit(playerHand) {
				return nil, ERR_TRIED_TO_SPLIT_BUT_NOT_ALLOWED
			}
			hand1, hand2, err := playerHand.SplitHand(bankrolledStrategy.Bankroll())
			if err != nil {
				return nil, err
			}
			hands1, err := PlayHand(hand1, dealerHand, bankrolledStrategy, rules, houseBankroll, shoe)
			if err != nil {
				return nil, err
			}
			hands2, err := PlayHand(hand2, dealerHand, bankrolledStrategy, rules, houseBankroll, shoe)
			if err != nil {
				return nil, err
			}
			return append(hands1, hands2...), err
		} else if action == STAND {
			break
		} else if action == HIT {
			card, err := shoe.Pop()
			if err != nil {
				return nil, err
			}
			playerHand.Push(card)
		} else {
			panic("I don't know this other thing to do ...")
		}
		if playerHand.Bust() {
			break
		}
	}
	return []Hand{playerHand}, nil
}

func SimulateSingleHand(shoeFactory ShoeFactory, handDealer HandDealer, dealerStrategy PlayStrategy, playerStrategy PlayStrategy, bettingStrategy BettingStrategy, number_of_iterations uint, rules Rules) (float64, error) {
	table := NewTable(NewDealer(dealerStrategy, handDealer), shoeFactory, uint(1), rules)
	player := NewPlayer(bettingStrategy, playerStrategy)
	table.SetPlayer(player, uint(0))
	for i := uint(0); i < number_of_iterations; i++ {
		err := table.PlayRound()
		if err != nil {
			return 0, err
		}
	}
	if player.Bankroll().CurrentBankroll()+table.Dealer().Bankroll().CurrentBankroll() != 0 {
		panic("Logic error: Bankrolls should even out")
	}
	return float64(bankroll.Money(player.Bankroll().CurrentBankroll()) / bankroll.Money(number_of_iterations)), nil
}

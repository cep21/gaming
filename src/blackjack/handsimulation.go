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

	for ; ; {
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

//
//func SimulateSingleHand(shoeFactory ShoeFactory, dealer HandDealer, dealerStrategy PlayStrategy, playerStrategy PlayStrategy, bettingStrategy BettingStrategy, number_of_iterations uint, rules Rules) (float64, error) {
//	playerBankroll := bankroll.NewMoneyHolder()
//	houseBankroll := bankroll.NewMoneyHolder()
//	//units_to_bet := bettingStrategy.GetMoneyToBet()
//	for i := uint(0); i < number_of_iterations; i++ {
//		shoe := shoeFactory.CreateShoe()
//		//bankroll.ChangeBankroll(-units_to_bet)
//		player_hands, dealerHand, err := dealer.DealHands(shoe, []BettingStrategy{bettingStrategy}, []bankroll.MoneyHolder{playerBankroll})
//		if err != nil {
//			return 0, err
//		}
//		playerHand := player_hands[0]
//		//playerHand := NewHand(player_hands[0], NewMoneyHolder(), 0)
//		//playerBankroll.TransferMoneyTo(playerHand.MoneyInThisHand(), units_to_bet)
//		if dealerHand.IsBlackjack() {
//			if playerHand.IsBlackjack() {
//				// Push the money back
//				playerHand.MoneyInThisHand().TransferMoneyTo(playerBankroll, playerHand.MoneyInThisHand().CurrentBankroll())
//				continue
//			} else {
//				playerHand.MoneyInThisHand().TransferMoneyTo(houseBankroll, playerHand.MoneyInThisHand().CurrentBankroll())
//				continue
//			}
//		} else if playerHand.IsBlackjack() {
//			houseBankroll.TransferMoneyTo(playerBankroll, playerHand.MoneyInThisHand().CurrentBankroll()*bankroll.Money(rules.BlackjackPayout()))
//			playerHand.MoneyInThisHand().TransferMoneyTo(playerBankroll, playerHand.MoneyInThisHand().CurrentBankroll())
//			continue
//		}
//
//		originalAllPlayerHands, err := PlayHand(playerHand, dealerHand, playerStrategy, rules, playerBankroll, houseBankroll, shoe)
//		playerHand = nil
//		if err != nil {
//			return 0, err
//		}
//		nonBustPlayerHands := []Hand{}
//		for _, hand := range originalAllPlayerHands {
//			if hand.Bust() {
//				hand.MoneyInThisHand().TransferMoneyTo(houseBankroll, hand.MoneyInThisHand().CurrentBankroll())
//			} else {
//				nonBustPlayerHands = append(nonBustPlayerHands, hand)
//			}
//		}
//		if len(nonBustPlayerHands) == 0 {
//			continue
//		}
//
//		allDealerHands, err := PlayHand(dealerHand, nil, dealerStrategy, rules, nil, nil, shoe)
//		if err != nil {
//			return 0, err
//		}
//		if len(allDealerHands) != 1 || allDealerHands[0] != dealerHand {
//			return 0, ERR_INVALID_DEALER_OPERATION
//		}
//
//		for _, finalPlayerHand := range nonBustPlayerHands {
//			if dealerHand.Bust() {
//				houseBankroll.TransferMoneyTo(playerBankroll, finalPlayerHand.MoneyInThisHand().CurrentBankroll())
//				finalPlayerHand.MoneyInThisHand().TransferMoneyTo(playerBankroll, finalPlayerHand.MoneyInThisHand().CurrentBankroll())
//				continue
//			}
//			//log.Printf("%s vs %s\n", player_hand, dealer_hand)
//			if dealerHand.Score() > finalPlayerHand.Score() {
//				// You loose, give it to the house
//				finalPlayerHand.MoneyInThisHand().TransferMoneyTo(houseBankroll, finalPlayerHand.MoneyInThisHand().CurrentBankroll())
//			} else if dealerHand.Score() < finalPlayerHand.Score() {
//				// You win.  Take some from the house
//				houseBankroll.TransferMoneyTo(playerBankroll, finalPlayerHand.MoneyInThisHand().CurrentBankroll())
//				finalPlayerHand.MoneyInThisHand().TransferMoneyTo(playerBankroll, finalPlayerHand.MoneyInThisHand().CurrentBankroll())
//			} else {
//				// Push the money back
//				finalPlayerHand.MoneyInThisHand().TransferMoneyTo(playerBankroll, finalPlayerHand.MoneyInThisHand().CurrentBankroll())
//			}
//		}
//	}
//	return float64(playerBankroll.CurrentBankroll())/float64(number_of_iterations), nil
//}

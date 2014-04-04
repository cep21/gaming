/**
 * Date: 4/2/14
 * Time: 2:02 PM
 * @author jack 
 */
package blackjack

func SimulateSingleHand(shoeFactory ShoeFactory, originalPlayerHand Hand, dealerCard Card, dealerStrategy PlayStrategy, bettingStrategy BettingStrategy, playerStrategy PlayStrategy, number_of_iterations uint) (float64, error) {
	sum_result := 0.0

	for i := uint(0); i < number_of_iterations; i++ {
		playerHand := originalPlayerHand.Clone()
		deck := shoeFactory.CreateShoe()
		units_to_bet := bettingStrategy.GetUnitsToBet()
		PlayHandOnStrategy(playerHand, dealerCard, playerStrategy, deck)
		if playerHand.Bust() {
			sum_result -= units_to_bet
			continue
		}

		dealer_hand := NewHand(dealerCard)
		card, err := deck.Pop()
		if err != nil {
			return 0, err
		}
		dealer_hand.Push(card)
		PlayHandOnStrategy(dealer_hand, /*Ignored*/dealerCard, dealerStrategy, deck)
		if dealer_hand.Bust() || playerHand.Score() > dealer_hand.Score() {
			sum_result += units_to_bet
			//fmt.Println("Win")
		} else if playerHand.Score() < dealer_hand.Score() {
			//fmt.Println("Lose")
			sum_result -= units_to_bet
		} else {
			// Push, zero
		}
	}
	return sum_result/float64(number_of_iterations), nil
}

//func SimulateSingleHand2(shoeFactory ShoeFactory, dealer HandDealer, dealerStrategy PlayStrategy,
//	playerStrategy PlayStrategy, bettingStrategy BettingStrategy,  number_of_iterations uint) (float64, error) {
//	sum_result := 0.0
//	bankroll := NewBankroll(0)
//	for i := uint(0); i < number_of_iterations; i++ {
//		//units_to_bet := bettingStrategy.GetUnitsToBet()
//		//shoe := shoeFactory.CreateShoe()
//		//bankroll.ChangeBankroll(-units_to_bet)
//		//hands := dealer.DealHands(shoe, 2)
//	}
//	return sum_result/float64(number_of_iterations), nil
//}

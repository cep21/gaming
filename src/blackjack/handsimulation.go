/**
 * Date: 4/2/14
 * Time: 2:02 PM
 * @author jack 
 */
package blackjack

func SimulateSingleHand(shoeFactory ShoeFactory, originalPlayerHand Hand, dealerCard Card, dealerStrategy PlayStrategy, bettingStrategy BettingStrategy, playerStrategy PlayStrategy, number_of_iterations uint) float64 {
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

		dealer_hand := NewHandWithCard(dealerCard)
		dealer_hand.Push(deck.Pop())
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
	return sum_result/float64(number_of_iterations)
}

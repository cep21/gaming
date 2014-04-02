/**
 * Date: 4/2/14
 * Time: 2:02 PM
 * @author jack 
 */
package blackjack

func SimulateSingleHand(cardsetFactory CardsetFactory, playerHand Hand, dealerCard Card, dealerStrategy ShouldHitStrategy, bettingStrategy BettingStrategy, playerStrategy ShouldHitStrategy, number_of_iterations uint) {
	sum_result := 0

	for i := uint(0); i < number_of_iterations; i++ {
		deck := cardsetFactory.CreateCardSet()
		units_to_bet := bettingStrategy.GetUnitsToBet()
		PlayHandOnStrategy(playerHand, dealerCard, playerStrategy, deck)
		if playerHand.Bust() {
			sum_result -= units_to_bet
		}

		dealer_hand := NewHandWithCard(dealerCard)
		dealer_hand.Push(deck.Pop())
		PlayHandOnStrategy(dealer_hand, /*Ignored*/dealerCard, dealerStrategy, deck)
		if dealer_hand.Bust() || playerHand.Score() > dealer_hand.Score() {
			sum_results += units_to_bet
		} else if playerHand.Score() < dealer_hand.Score() {
			sum_results -= units_to_bet
		} else {
			// Push
		}
	}
}

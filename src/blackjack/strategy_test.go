/**
 * Date: 4/1/14
 * Time: 3:49 PM
 * @author jack 
 */
package blackjack

import (
	"testing"
	"gaming"
)

func TestDealerSoftHit(t *testing.T) {
	strat := NewDealerStrategy(true)
	hand := NewHand()
	hand.Push(NewCard(gaming.Spade, Ace))
	hand.Push(NewCard(gaming.Spade, Six))
	if !strat.ShouldHit(hand, hand) {
		t.Error("Expect to hit on soft 17")
	}
	hand.Push(NewCard(gaming.Spade, Two))
	if strat.ShouldHit(hand, hand) {
		t.Error("Expect to not hit")
	}
}

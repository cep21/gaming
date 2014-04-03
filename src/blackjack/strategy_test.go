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
	if strat.TakeAction(hand, NewCard(gaming.Spade, Six)) != HIT {
		t.Error("Expect to hit on soft 17")
	}
	hand.Push(NewCard(gaming.Spade, Two))
	if strat.TakeAction(hand, NewCard(gaming.Spade, Six)) != STAND {
		t.Error("Expect to not hit")
	}
}

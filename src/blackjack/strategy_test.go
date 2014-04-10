/**
 * Date: 4/1/14
 * Time: 3:49 PM
 * @author jack
 */
package blackjack

import (
	"gaming"
	"testing"
	"math/rand"
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

func TestLearnAction(t *testing.T) {
	rules := NewRulesetFactory().SurrenderOption(LATE_SURRENDER).Build()
	//	rules := NewRulesetFactory().Build()
	//rules Rules, shoeFactory ShoeFactory, handDealer HandDealer, dealerStrategy PlayStrategy, iterations uint
	r := rand.New(rand.NewSource(3))
	shoeFactory := NewClonedDeckFactory(NewInfiniteShoe(r), r)
	dealerStrategy := NewDealerStrategy(rules.DealerHitOnSoft17())
	iterations := uint(10000)
	strat := NewDiscoveredStrategy(rules, shoeFactory, dealerStrategy, iterations)
	verifyAction(t, NewHand(NewCard(gaming.Spade, Two), NewCard(gaming.Spade, Two), NewCard(gaming.Spade, Six)), NewCard(gaming.Heart, Four), HIT, strat)
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Ten)), NewCard(gaming.Heart, Nine), STAND, strat)
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Ten)), NewCard(gaming.Heart, Eight), STAND, strat)
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Ten)), NewCard(gaming.Heart, Seven), STAND, strat)
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Ten)), NewCard(gaming.Heart, Six), STAND, strat)
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Six)), NewCard(gaming.Heart, Ten), SURRENDER, strat)
//
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Ten)), NewCard(gaming.Heart, Six), STAND, strat)
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Eight), NewCard(gaming.Spade, Eight)), NewCard(gaming.Heart, Seven), SPLIT, strat)
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Eight), NewCard(gaming.Spade, Three)), NewCard(gaming.Heart, Ten), DOUBLE, strat)
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Three)), NewCard(gaming.Heart, Three), STAND, strat)
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Two)), NewCard(gaming.Heart, Three), HIT, strat)
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Seven)), NewCard(gaming.Heart, Nine), STAND, strat)
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Six)), NewCard(gaming.Heart, Nine), SURRENDER, strat)
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Ace), NewCard(gaming.Spade, Five)), NewCard(gaming.Heart, Nine), HIT, strat)
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Ten)), NewCard(gaming.Heart, Eight), STAND, strat)
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Ten)), NewCard(gaming.Heart, Seven), STAND, strat)
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Ten)), NewCard(gaming.Heart, Six), STAND, strat)
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Five), NewCard(gaming.Spade, Ten)), NewCard(gaming.Heart, Ten), SURRENDER, strat)
//	verifyAction(t, NewHand(NewCard(gaming.Spade, Five), NewCard(gaming.Spade, Ten)), NewCard(gaming.Heart, Six), STAND, strat)
}

func verifyAction(t *testing.T, userHand Hand, dealerCard Card, expectedBestAction GameAction, strat *DiscoveredStrategy) {
	action := strat.TakeAction(userHand, dealerCard)
	if action != expectedBestAction {
		t.Errorf("I expected it to be best to %s on %s vs %s, but it's better to %s\n", expectedBestAction, userHand, dealerCard, action)
	}
}

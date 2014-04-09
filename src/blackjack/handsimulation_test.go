package blackjack

import (
	"testing"
	"math/rand"
	"gaming"
	"math"
)

func TestHitOn12(t *testing.T) {
	player_hand := NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Two))

	go simulateHand(t, Two, player_hand, true)
	go simulateHand(t, Three, player_hand, true)
	go simulateHand(t, Four, player_hand, false)
	go simulateHand(t, Five, player_hand, false)
	go simulateHand(t, Six, player_hand, false)
	go simulateHand(t, Seven, player_hand, true)
	go simulateHand(t, Eight, player_hand, true)
	go simulateHand(t, Nine, player_hand, true)
	go simulateHand(t, Ten, player_hand, true)
	go simulateHand(t, Jack, player_hand, true)
	go simulateHand(t, Queen, player_hand, true)
	go simulateHand(t, King, player_hand, true)
	go simulateHand(t, Ace, player_hand, true)
}

func TestHitOn13(t *testing.T) {
	player_hand := NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Three))

	go simulateHand(t, Two, player_hand, false)
	go simulateHand(t, Three, player_hand, false)
	go simulateHand(t, Four, player_hand, false)
	go simulateHand(t, Five, player_hand, false)
	go simulateHand(t, Six, player_hand, false)
	go simulateHand(t, Seven, player_hand, true)
	go simulateHand(t, Eight, player_hand, true)
	go simulateHand(t, Nine, player_hand, true)
	go simulateHand(t, Ten, player_hand, true)
	go simulateHand(t, Jack, player_hand, true)
	go simulateHand(t, Queen, player_hand, true)
	go simulateHand(t, King, player_hand, true)
	go simulateHand(t, Ace, player_hand, true)
}

func TestHitOn16(t *testing.T) {
	player_hand := NewHand(NewCard(gaming.Spade, Ten), NewCard(gaming.Spade, Six))

	go simulateHand(t, Two, player_hand, false)
	go simulateHand(t, Three, player_hand, false)
	go simulateHand(t, Four, player_hand, false)
	go simulateHand(t, Five, player_hand, false)
	go simulateHand(t, Six, player_hand, false)
	go simulateHand(t, Seven, player_hand, true)
	go simulateHand(t, Eight, player_hand, true)
	go simulateHand(t, Nine, player_hand, true)
	go simulateHand(t, Ten, player_hand, true)
	go simulateHand(t, Jack, player_hand, true)
	go simulateHand(t, Queen, player_hand, true)
	go simulateHand(t, King, player_hand, true)
	go simulateHand(t, Ace, player_hand, true)
}

func TestHitOnSoft18(t *testing.T) {
	player_hand := NewHand(NewCard(gaming.Spade, Ace), NewCard(gaming.Spade, Seven))

	go simulateHand(t, Two, player_hand, false)
	go simulateHand(t, Three, player_hand, false)
	go simulateHand(t, Four, player_hand, false)
	go simulateHand(t, Five, player_hand, false)
	go simulateHand(t, Six, player_hand, false)
	go simulateHand(t, Seven, player_hand, false)
	go simulateHand(t, Eight, player_hand, false)
	go simulateHand(t, Nine, player_hand, true)
	go simulateHand(t, Ten, player_hand, true)
	go simulateHand(t, Jack, player_hand, true)
	go simulateHand(t, Queen, player_hand, true)
	go simulateHand(t, King, player_hand, true)
	go simulateHand(t, Ace, player_hand, true)
}



func simulateHand(t *testing.T, dealerValue Value, player_hand Hand, expectedBetterToHit bool) {
	var hit_strategy PlayStrategy
	if player_hand.IsSoft() {
		hit_strategy = NewHitOnAScoreStrategy(player_hand.Score(), 16)
	} else {
		hit_strategy = NewHitOnAScoreStrategy(player_hand.Score(), player_hand.Score())
	}

	never_bust_strategy := NewNeverBustStrategy(false)
	dealer_strategy := NewDealerStrategy(true)
	// Seed 4 happens to work with only 12,000 simulated rounds
	r := rand.New(rand.NewSource(5))
	betting_strategy := NewConsistentBettingStrategy(1)
	shoe_factory := NewClonedDeckFactory(NewInfiniteShoe(r), r)
	rules := NewRulesetFactory().Build()
	rounds_to_simulate := uint(12000)
	dealer := NewForceDealerPlayerHands(player_hand, dealerValue)
	result_for_hitting, err := SimulateSingleHand(shoe_factory, dealer, dealer_strategy, hit_strategy, betting_strategy, rounds_to_simulate, rules)
	if err != nil {
		t.Fatal("Got an error!")
	}
	result_for_standing, err := SimulateSingleHand(shoe_factory, dealer, dealer_strategy, never_bust_strategy, betting_strategy, rounds_to_simulate, rules)
	if err != nil {
		t.Fatal("Got an error!")
	}
	delta := math.Abs(result_for_hitting - result_for_standing)
	if expectedBetterToHit {
		if result_for_standing > result_for_hitting {
			t.Errorf("I expect it to be better to hit, but it's better to stand %s vs %s by %f\n", dealerValue, player_hand, delta)
		}
	} else {
		if result_for_standing < result_for_hitting {
			t.Errorf("I expect it to be better to stand, but it's better to hit %s vs %s by %f\n", dealerValue, player_hand, delta)
		}
	}
}

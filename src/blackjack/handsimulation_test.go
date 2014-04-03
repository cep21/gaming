package blackjack

import (
	"testing"
	"math/rand"
	"gaming"
	"math"
)

func TestHitOn12(t *testing.T) {
	player_hand := NewHand()
	player_hand.Push(NewCard(gaming.Spade, Ten))
	player_hand.Push(NewCard(gaming.Spade, Two))

	simulateHand(t, Two, player_hand,    true)
	simulateHand(t, Three, player_hand,  true)
	simulateHand(t, Four, player_hand,   false)
	simulateHand(t, Five, player_hand,   false)
	simulateHand(t, Six, player_hand,    false)
	simulateHand(t, Seven, player_hand,  true)
	simulateHand(t, Eight, player_hand,  true)
	simulateHand(t, Nine, player_hand,   true)
	simulateHand(t, Ten, player_hand,    true)
	simulateHand(t, Jack, player_hand,   true)
	simulateHand(t, Queen, player_hand,  true)
	simulateHand(t, King, player_hand,   true)
}

func TestHitOn13(t *testing.T) {
	player_hand := NewHand()
	player_hand.Push(NewCard(gaming.Spade, Ten))
	player_hand.Push(NewCard(gaming.Spade, Three))

	simulateHand(t, Two, player_hand,    false)
	simulateHand(t, Three, player_hand,  false)
	simulateHand(t, Four, player_hand,   false)
	simulateHand(t, Five, player_hand,   false)
	simulateHand(t, Six, player_hand,    false)
	simulateHand(t, Seven, player_hand,  true)
	simulateHand(t, Eight, player_hand,  true)
	simulateHand(t, Nine, player_hand,   true)
	simulateHand(t, Ten, player_hand,    true)
	simulateHand(t, Jack, player_hand,   true)
	simulateHand(t, Queen, player_hand,  true)
	simulateHand(t, King, player_hand,   true)
}



func simulateHand(t *testing.T, dealerValue Value, player_hand Hand, expectedBetterToHit bool) {
	hit_strategy := NewHitOnAScoreStrategy(player_hand.Score())
	never_bust_strategy := NewNeverBustStrategy(false)
	dealer_strategy := NewDealerStrategy(true)
	r := rand.New(rand.NewSource(2))
	betting_strategy := NewConsistentBettingStrategy(1)
	infinite_deck := NewClonedDeckFactory(NewInfiniteShoe(r), r)
	dealerCard := NewCard(gaming.Spade, dealerValue)
	rounds_to_simulate := uint(50000)
	result_for_hitting := SimulateSingleHand(infinite_deck, player_hand, dealerCard, dealer_strategy, betting_strategy, hit_strategy, rounds_to_simulate)
	result_for_standing := SimulateSingleHand(infinite_deck, player_hand, dealerCard, dealer_strategy, betting_strategy, never_bust_strategy, rounds_to_simulate)
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

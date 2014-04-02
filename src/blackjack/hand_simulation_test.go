package blackjack

import (
	"testing"
)

func TestHitOn12(t *testing.T) {
	hit_on_12 := NewHitOnAScoreStrategy(12)
	dealer_strategy := NewDealerStrategy(true)

}

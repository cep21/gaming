package blackjack

import (
	"testing"
	"gaming"
)

func TestSuits(t *testing.T) {
	if Ace.Score() != 11 {
		t.Error("Ace should have a score of 11")
	}

	if Ace.Name() != "ace" {
		t.Error("Ace should be named ace")
	}

	if len(gaming.Suits()) != 4 {
		t.Error("Expect 4 suits")
	}
}

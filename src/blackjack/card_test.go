package blackjack

import (
	"testing"
	"gaming"
	"math/rand"
)

func TestSuits(t *testing.T) {
	if Ace.Score() != 1 {
		t.Error("Ace should have a score of 1")
	}

	if Ace.Name() != "ace" {
		t.Error("Ace should be named ace")
	}

	if len(gaming.Suits()) != 4 {
		t.Error("Expect 4 suits")
	}
}

func TestRandomCard(t *testing.T) {
	r := rand.New(rand.NewSource(2))
	c1 := NewRandomCard(r)
	c2 := NewRandomCard(r)
	if c1.Value() == c2.Value() {
		t.Error("Did not expect the same value (note constant seed)")
	}
}

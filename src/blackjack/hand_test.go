package blackjack

import (
	"testing"
	"gaming"
)

func TestSplit(t *testing.T) {
	h := NewHand()
	h.Push(NewCard(gaming.Spade, Jack))
	h.Push(NewCard(gaming.Spade, Ten))
	if !h.CanSplit() {
		t.Error("Expected to be able to split")
	}
	h.Push(NewCard(gaming.Spade, Ace))
	if h.CanSplit() {
		t.Error("Expected to not be able to split")
	}

	h = NewHand()
	h.Push(NewCard(gaming.Spade, Jack))
	h.Push(NewCard(gaming.Spade, Two))
	if h.CanSplit() {
		t.Error("Expected not to be able to split")
	}
}

func TestHand(t *testing.T) {
	h := NewHand()
	h.Push(NewCard(gaming.Spade, Ace))
	h.Push(NewCard(gaming.Spade, Ten))
	if h.Score() != 21 {
		t.Error("Expected score of 21")
	}
	if !h.IsSoft() {
		t.Error("Expected Soft hand")
	}
	if !h.IsBlackjack() {
		t.Error("Expected blackjack hand")
	}
	h.Push(NewCard(gaming.Spade, Six))
	if h.Score() != 17 {
		t.Errorf("Expected score of 17, not %d\n", h.Score())
	}
	if h.IsSoft() {
		t.Error("Expected non Soft hand")
	}
	h.Push(NewCard(gaming.Spade, Ace))
	if h.Score() != 18 {
		t.Errorf("Expected score of 18, not %d\n", h.Score())
	}
	if h.IsSoft() {
		t.Error("Expected non Soft hand")
	}
	h.Push(NewCard(gaming.Spade, Ace))
	h.Push(NewCard(gaming.Spade, Ace))
	h.Push(NewCard(gaming.Spade, Ace))
	h.Push(NewCard(gaming.Spade, Ace))
	if h.Score() != 22 {
		t.Errorf("Expected score of 22, not %d\n", h.Score())
	}
}

func TestSoft(t *testing.T) {
	h := NewHand()
	h.Push(NewCard(gaming.Spade, Ace))
	h.Push(NewCard(gaming.Spade, Six))
	if h.Score() != 17 {
		t.Error("Expected score of 17")
	}
	if !h.IsSoft() {
		t.Error("Expected Soft hand")
	}
	h.Push(NewCard(gaming.Spade, Two))
	if h.Score() != 19 {
		t.Errorf("Expected score of 19, not %d\n", h.Score())
	}
	if !h.IsSoft() {
		t.Error("Expected Soft hand")
	}

	h.Push(NewCard(gaming.Spade, Ace))
	if h.Score() != 20 {
		t.Errorf("Expected score of 20, not %d\n", h.Score())
	}
	if !h.IsSoft() {
		t.Error("Expected soft hand")
	}
	h.Push(NewCard(gaming.Spade, Six))
	if h.Score() != 16 {
		t.Errorf("Expected score of 16, not %d\n", h.Score())
	}
	if h.IsSoft() {
		t.Error("Expected non soft hand")
	}
}

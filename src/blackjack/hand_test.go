package blackjack

import (
	"gaming"
	"testing"
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

func TestSoftAgain(t *testing.T) {
	// 5 4 => 9        (Hard)
	// 5 4 A => 20     (Soft)
	// 5 4 A A => 21   (Soft)
	// 5 4 A A 2 => 13 (Hard)
	// 5 4 A A 2 8 => 21 (Hard)
	// 5 4 A A 2 8 A => 22 (Hard)
	type handScore struct {
		value         Value
		expectedScore uint
		isSoft        bool
	}
	v := []handScore{
		handScore{Five, 5, false},
		handScore{Four, 9, false},
		handScore{Ace, 20, true},
		handScore{Ace, 21, true},
		handScore{Two, 13, false},
		handScore{Eight, 21, false},
		handScore{Ace, 22, false},
	}
	h := NewHand()
	for _, s := range v {
		h.Push(NewCard(gaming.Spade, s.value))
		if h.Score() != s.expectedScore {
			t.Error("Did not get expected value")
		}
		if h.IsSoft() != s.isSoft {
			t.Error("Softness not as expected")
		}
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
	if h.Bust() {
		t.Error("Expected not to bust")
	}
	h.Push(NewCard(gaming.Spade, Ace))
	if h.Score() != 22 {
		t.Errorf("Expected score of 22, not %d\n", h.Score())
	}
	if !h.Bust() {
		t.Error("Expected to bust")
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

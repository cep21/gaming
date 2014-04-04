package blackjack

type BlackjackPayout float64

const NORMAL_PAYOUT = 1.5
const SIXFIVE_PAYOUT = 6.0/5.0

type Rules interface {
	DealerHitOnSoft17() bool
	DoubleAfterSplit() bool
	ResplitAces() bool
	CanDouble(hand Hand) bool
	CanSplit(hand Hand) bool
	BlackjackPayout() BlackjackPayout
	ReshufflePenetration() float64
}

type rulesImpl struct {
	dealerHitOnSoft17     bool
	doubleAfterSplit      bool

	resplitAces           bool
	hardDoubleValues      []uint
	hardDoubleLookup      []bool
	softDoubleValues      []uint
	softDoubleLookup      []bool
	blackjackPayout       BlackjackPayout
	reshufflePenetration  float64
}

func (this *rulesImpl) DealerHitOnSoft17() bool {
	return this.dealerHitOnSoft17
}

func (this *rulesImpl) DoubleAfterSplit() bool {
	return this.doubleAfterSplit
}

func (this *rulesImpl) ResplitAces() bool {
	return this.resplitAces
}

func (this *rulesImpl) BlackjackPayout() BlackjackPayout {
	return this.blackjackPayout
}

func (this *rulesImpl) ReshufflePenetration() float64 {
	return this.reshufflePenetration
}

func (this *rulesImpl) CanSplit(hand Hand) bool {
	return hand.CanSplit()
}

func (this *rulesImpl) CanDouble(hand Hand) bool {
	if hand.IsSoft() {
		return this.softDoubleLookup[hand.Score()]
	} else {
		return this.hardDoubleLookup[hand.Score()]
	}
}

func NewRuleset(softhit bool, doublesplit bool, resplitace bool, hardDoubleValues[]uint, softDoubleValues []uint, payout BlackjackPayout, penetration float64) Rules {
	hardDoubles := make([]bool, 21)
	for _, v := range hardDoubleValues {
		hardDoubles[v] = true
	}

	softDoubles := make([]bool, 21)
	for _, v := range softDoubleValues {
		softDoubles[v] = true
	}
	return &rulesImpl{softhit, doublesplit, resplitace, hardDoubleValues, hardDoubles, softDoubleValues, softDoubles, payout, penetration}
}

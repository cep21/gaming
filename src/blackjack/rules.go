package blackjack

type BlackjackPayout float64

const NORMAL_PAYOUT = 1.5
const SIXFIVE_PAYOUT = 6.0 / 5.0

type Rules interface {
	DealerHitOnSoft17() bool
	DoubleAfterSplit() bool
	ResplitAces() bool
	CanDouble(handValue uint) bool
	BlackjackPayout() BlackjackPayout
	ReshufflePenetration() float64
}

type rulesImpl struct {
	dealerHitOnSoft17 bool
	doubleAfterSplit  bool
	resplitAces       bool
	doubleValues      [] uint
	doubleLookup      [] bool
	blackjackPayout   BlackjackPayout
	reshufflePenetration float64
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


func (this *rulesImpl) CanDouble(handValue uint) bool {
	return this.doubleLookup[handValue]
}

func NewRuleset(softhit bool, doublesplit bool, resplitace bool, doubleval[]uint, payout BlackjackPayout, penetration float64) Rules {
	doubleLookup := make([]bool, 21)
	for _, v := range doubleval {
		doubleLookup[v] = true
	}
	return &rulesImpl{softhit, doublesplit, resplitace, doubleval, doubleLookup, payout, penetration}
}

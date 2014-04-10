package blackjack

func AllValues() []uint {
	return []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
}

type BlackjackPayout float64

const NORMAL_PAYOUT = 1.5
const SIXFIVE_PAYOUT = 6.0 / 5.0

type SurrenderOption interface {
	Name() string
}

type surrenderOptionImpl struct {
	name string
}

func (this *surrenderOptionImpl) Name() string {
	return this.name
}

func (this *surrenderOptionImpl) String() string {
	return this.name
}

var NO_SURRENDER = &surrenderOptionImpl{"none"}
var EARLY_SURRENDER = &surrenderOptionImpl{"early"}
var LATE_SURRENDER = &surrenderOptionImpl{"late"}

type Rules interface {
	DealerHitOnSoft17() bool
	DoubleAfterSplit() bool
	MaxResplits() uint
	ResplitAces() bool
	CanDouble(hand Hand) bool
	CanSplit(hand Hand) bool
	CanHit(hand Hand) bool
	CanSurrender(hand Hand) bool
	AllowsAction(GameAction, Hand) bool
	BlackjackPayout() BlackjackPayout
	ReshufflePenetration() float64
}

type rulesImpl struct {
	doubleAfterSplit bool
	maxResplits      uint
	hitSoft17        bool
	resplitAces      bool
	hardDoubleValues []uint
	softDoubleValues []uint
	payout           BlackjackPayout
	penetration      float64
	surrenderOption  SurrenderOption

	hardDoubleLookup []bool
	softDoubleLookup []bool
}

func (this *rulesImpl) DealerHitOnSoft17() bool {
	return this.hitSoft17
}

func (this *rulesImpl) AllowsAction(action GameAction, playerHand Hand) bool {
	if action == HIT {
		return this.CanHit(playerHand)
	} else if action == STAND {
		return true
	} else if action == DOUBLE {
		return this.CanDouble(playerHand)
	} else if action == SPLIT {
		return this.CanSplit(playerHand)
	} else if action == SURRENDER {
		return this.CanSurrender(playerHand)
	} else {
		panic("Unknown action!")
	}
}


func (this *rulesImpl) DoubleAfterSplit() bool {
	return this.doubleAfterSplit
}

func (this *rulesImpl) MaxResplits() uint {
	return this.maxResplits
}

func (this *rulesImpl) ResplitAces() bool {
	return this.resplitAces
}

func (this *rulesImpl) BlackjackPayout() BlackjackPayout {
	return this.payout
}

func (this *rulesImpl) ReshufflePenetration() float64 {
	return this.penetration
}

func (this *rulesImpl) CanSplit(hand Hand) bool {
	if hand.IsSplitHand() {
		if hand.FirstCard().Value() == Ace {
			return this.ResplitAces()
		} else {
			return hand.CanSplit()
		}
	} else {
		return hand.CanSplit()
	}
}

func (this *rulesImpl) CanSurrender(hand Hand) bool {
	if hand.Size() > 2 {
		return false
	} else if hand.Size() == 2 {
		return this.surrenderOption != NO_SURRENDER
	} else {
		return this.surrenderOption == EARLY_SURRENDER
	}
}

func (this *rulesImpl) CanHit(hand Hand) bool {
	if hand.Bust() {
		return false
	}
	if hand.IsSplitHand() {

	}
	return true
}

func (this *rulesImpl) CanDouble(hand Hand) bool {
	if len(hand.Cards()) !=2 {
		return false;
	}
	if hand.IsSplitHand() {
		if !this.doubleAfterSplit {
			return false
		}
	}
	if hand.IsSoft() {
		return this.softDoubleLookup[hand.Score()]
	} else {
		return this.hardDoubleLookup[hand.Score()]
	}
}

type RulesetFactory interface {
	Build() Rules
	DoubleAfterSplit(doubleAfterSplit bool) RulesetFactory
	MaxResplits(maxResplits uint) RulesetFactory
	DealerHitsSoft17(hitSoft17 bool) RulesetFactory
	ResplitAces(resplitAces bool) RulesetFactory
	HardDoubleValues(hardDoubleValues []uint) RulesetFactory
	SoftDoubleValues(softDoubleValues []uint) RulesetFactory
	Payout(payout BlackjackPayout) RulesetFactory
	Penetration(penetration float64) RulesetFactory
	SurrenderOption(surrenderOption SurrenderOption) RulesetFactory
}

type rulesetFactoryImpl struct {
	doubleAfterSplit bool
	maxResplits      uint
	hitSoft17        bool
	resplitAces      bool
	hardDoubleValues []uint
	softDoubleValues []uint
	payout           BlackjackPayout
	penetration      float64
	surrenderOption  SurrenderOption
}

var DEFAULT_DOUBLE_AFTER_SPLIT = true
var DEFAULT_RESPLIT_ACES = true
var DEFAULT_MAX_RESPLITS = uint(4)
var DEFAULT_HIT_SOFT_17 = true
var DEFAULT_HARD_DOUBLE_VALUES = AllValues()
var DEFAULT_SOFT_DOUBLE_VALUES = AllValues()
var DEFAULT_PAYOUT = NORMAL_PAYOUT
var DEFAULT_PENETRATION = .8
var DEFAULT_SURRENDER_OPTION = NO_SURRENDER

func NewRulesetFactory() RulesetFactory {
	return &rulesetFactoryImpl{
		doubleAfterSplit: DEFAULT_DOUBLE_AFTER_SPLIT,
		maxResplits:      DEFAULT_MAX_RESPLITS,
		hitSoft17:        DEFAULT_HIT_SOFT_17,
		resplitAces:      DEFAULT_RESPLIT_ACES,
		hardDoubleValues: DEFAULT_HARD_DOUBLE_VALUES,
		softDoubleValues: DEFAULT_SOFT_DOUBLE_VALUES,
		payout:           BlackjackPayout(DEFAULT_PAYOUT),
		penetration:      DEFAULT_PENETRATION,
		surrenderOption:  DEFAULT_SURRENDER_OPTION,
	}
}

func (this *rulesetFactoryImpl) Build() Rules {
	hardDoubles := make([]bool, 22)
	for _, v := range this.hardDoubleValues {
		hardDoubles[v] = true
	}

	softDoubles := make([]bool, 22)
	for _, v := range this.softDoubleValues {
		softDoubles[v] = true
	}

	return &rulesImpl{
		doubleAfterSplit: this.doubleAfterSplit,
		maxResplits:      this.maxResplits,
		hitSoft17:        this.hitSoft17,
		resplitAces:      this.resplitAces,
		hardDoubleValues: this.hardDoubleValues,
		softDoubleValues: this.softDoubleValues,
		payout:           this.payout,
		penetration:      this.penetration,
		surrenderOption:  this.surrenderOption,
		hardDoubleLookup: hardDoubles,
		softDoubleLookup: softDoubles,
	}
}

func (this *rulesetFactoryImpl) DoubleAfterSplit(doubleAfterSplit bool) RulesetFactory {
	this.doubleAfterSplit = doubleAfterSplit
	return this
}
func (this *rulesetFactoryImpl) DealerHitsSoft17(hitSoft17 bool) RulesetFactory {
	this.hitSoft17 = hitSoft17
	return this
}
func (this *rulesetFactoryImpl) MaxResplits(maxResplits uint) RulesetFactory {
	this.maxResplits = maxResplits
	return this
}
func (this *rulesetFactoryImpl) ResplitAces(resplitAces bool) RulesetFactory {
	this.resplitAces = resplitAces
	return this
}
func (this *rulesetFactoryImpl) HardDoubleValues(hardDoubleValues []uint) RulesetFactory {
	this.hardDoubleValues = hardDoubleValues
	return this
}
func (this *rulesetFactoryImpl) SoftDoubleValues(softDoubleValues []uint) RulesetFactory {
	this.softDoubleValues = softDoubleValues
	return this
}
func (this *rulesetFactoryImpl) Payout(payout BlackjackPayout) RulesetFactory {
	this.payout = payout
	return this
}
func (this *rulesetFactoryImpl) Penetration(penetration float64) RulesetFactory {
	this.penetration = penetration
	return this
}
func (this *rulesetFactoryImpl) SurrenderOption(surrenderOption SurrenderOption) RulesetFactory {
	this.surrenderOption = surrenderOption
	return this
}

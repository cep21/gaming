/**
 * Date: 4/1/14
 * Time: 3:46 PM
 * @author jack
 */
package blackjack

import (
	"math"
	//	"log"
	"fmt"
	"io"
)

type GameAction interface {
	Name() string
	Symbol() rune
	Index() int
}

type gameActionImpl struct {
	name   string
	symbol rune
	index  int
}

func (this *gameActionImpl) Name() string {
	return this.name
}

func (this *gameActionImpl) Symbol() rune {
	return this.symbol
}

func (this *gameActionImpl) String() string {
	return this.Name()
}

func (this *gameActionImpl) Index() int {
	return this.index
}

var HIT = &gameActionImpl{"hit", 'H', 0}
var STAND = &gameActionImpl{"stand", 'S', 1}
var DOUBLE = &gameActionImpl{"double", 'D', 2}
var SPLIT = &gameActionImpl{"split", 'P', 3}
var SURRENDER = &gameActionImpl{"surrender", 'U', 4}

type strategyTableAction interface {
	Name() string
	ShortName() string
	Index() int
	GameAction(rules Rules, hand Hand) GameAction
	Combine(GameAction) strategyTableAction
}

type strategyTableActionImpl struct {
	name      string
	shortName string
	index     int
}

var stratHit = &strategyTableActionImpl{"hit", "H", 0}
var stratStand = &strategyTableActionImpl{"stand", "S", 1}
var stratSplit = &strategyTableActionImpl{"split", "P", 2}
var stratDoubleUnknown = &strategyTableActionImpl{"double", "D", 3}
var stratDoubleHit = &strategyTableActionImpl{"double/hit", "Dh", 4}
var stratDoubleStand = &strategyTableActionImpl{"double/stand", "Ds", 5}
var stratSurrenderUnknown = &strategyTableActionImpl{"sur", "U", 6}
var stratSurrenderStand = &strategyTableActionImpl{"sur/stand", "Us", 7}
var stratSurrenderHit = &strategyTableActionImpl{"sur/hit", "Uh", 8}

func combineTableAction(stratAction strategyTableAction, gameAction GameAction) strategyTableAction {
	if stratAction == nil {
		return getStrategyTableAction(gameAction)
	} else {
		return stratAction.Combine(gameAction)
	}
}

func getStrategyTableAction(gameAction GameAction) strategyTableAction {
	if gameAction == HIT {
		return stratHit
	} else if gameAction == STAND {
		return stratStand
	} else if gameAction == SURRENDER {
		return stratSurrenderUnknown
	} else if gameAction == DOUBLE {
		return stratDoubleUnknown
	} else if gameAction == SPLIT {
		return stratSplit
	}
	panic("Unknown game action!")
}

func (this *strategyTableActionImpl) Name() string {
	return this.name
}

func (this *strategyTableActionImpl) ShortName() string {
	return this.shortName
}

func (this *strategyTableActionImpl) String() string {
	return this.Name()
}

func (this *strategyTableActionImpl) Index() int {
	return this.index
}

func (this *strategyTableActionImpl) Combine(gameAction GameAction) strategyTableAction {
	switch this.ShortName(){
	case "D":
		if gameAction == HIT {
			return stratDoubleHit
		} else if gameAction == STAND {
			return stratDoubleStand
		}
		break
	case "U":
		if gameAction == HIT {
			return stratSurrenderHit
		} else if gameAction == STAND {
			return stratSurrenderStand
		}
		break
	}

	panic(fmt.Sprintf("Unknown how to combine these two %s and %s!!!??", *this, gameAction))
}

func (this *strategyTableActionImpl) GameAction(rules Rules, hand Hand) GameAction {
	switch this.ShortName() {
	case "H":
		return HIT
	case "S":
		return STAND
	case "P":
		if rules.CanSplit(hand) {
			return SPLIT
		} else {
			panic("This should never happen!")
		}
	case "D":
		if rules.CanDouble(hand) {
			return DOUBLE
		} else {
			return nil
		}
	case "Dh":
		if rules.CanDouble(hand) {
			return DOUBLE
		} else {
			return HIT
		}
	case "Ds":
		if rules.CanDouble(hand) {
			return DOUBLE
		} else {
			return STAND
		}
	case "U":
		if rules.CanSurrender(hand) {
			return SURRENDER
		} else {
			return nil
		}
	case "Us":
		if rules.CanSurrender(hand) {
			return SURRENDER
		} else {
			return STAND
		}
	case "Uh":
		if rules.CanSurrender(hand) {
			return SURRENDER
		} else {
			return HIT
		}
	default:
		panic("Unknown symbol logic error")
	}
	panic("Unknown symbol logic error")
}

type PlayStrategy interface {
	TakeAction(currentHand Hand, shownCard Card) GameAction
}

type dealerHitStrategy struct {
	hitOnSoft17 bool
}

func (this *dealerHitStrategy) TakeAction(currentHand Hand, _ Card) GameAction {
	score := currentHand.Score()
	if score == 17 {
		if this.hitOnSoft17 && currentHand.IsSoft() {
			return HIT
		} else {
			return STAND
		}
	} else {
		if score < 17 {
			return HIT
		} else {
			return STAND
		}
	}
}

func NewDealerStrategy(hitOnSoft17 bool) PlayStrategy {
	return &dealerHitStrategy{hitOnSoft17: hitOnSoft17}
}

// Always hit
type alwaysHitStrategy struct {
}

func (this *alwaysHitStrategy) TakeAction(_ Hand, _ Card) GameAction {
	return HIT
}

func NewAlwaysHitStrategy() PlayStrategy {
	return &alwaysHitStrategy{}
}

// Always stand
type alwaysStandStrategy struct {
}

func (this *alwaysStandStrategy) TakeAction(_ Hand, _ Card) GameAction {
	return STAND
}

func NewAlwaysStandStrategy() PlayStrategy {
	return &alwaysStandStrategy{}
}

// Never bust
type neverBustStrategy struct {
	shouldHitSoft bool
}

func (this *neverBustStrategy) TakeAction(currentHand Hand, _ Card) GameAction {
	if (currentHand.IsSoft() && this.shouldHitSoft) || currentHand.Score() < 12 {
		return HIT
	} else {
		return STAND
	}
}

func NewNeverBustStrategy(should_hit_soft bool) PlayStrategy {
	return &neverBustStrategy{should_hit_soft}
}

func PlayHandOnStrategy(currentHand Hand, shownCard Card, strategy PlayStrategy, deck Shoe) {
	for deck.CardsLeft() != 0 {
		if currentHand.Bust() {
			return
		}
		res := strategy.TakeAction(currentHand, shownCard)
		if res == STAND {
			return
		} else if res == HIT {
			c, err := deck.Pop()
			if err != nil {
				return // err
			}
			currentHand.Push(c)
		} else {
			panic("I don't have SPLIT or DOUBLE yet")
		}
	}
}

// Hit on a hard score strategy
type hitOnAScoreStrategy struct {
	hardScoreToHit uint
	softScoreToHit uint
}

func (this *hitOnAScoreStrategy) TakeAction(currentHand Hand, _ Card) GameAction {
	if (currentHand.IsSoft() && currentHand.Score() <= this.softScoreToHit) || (!currentHand.IsSoft() && currentHand.Score() <= this.hardScoreToHit) {
		return HIT
	} else {
		return STAND
	}
}

func NewHitOnAScoreStrategy(soft_score_to_hit uint, hard_score_to_hit uint) PlayStrategy {
	return &hitOnAScoreStrategy{hard_score_to_hit, soft_score_to_hit}
}

type DiscoveredStrategy struct {
	rules           Rules
	shoeFactory     ShoeFactory
	dealerStrategy  PlayStrategy
	bettingStrategy BettingStrategy
	iterations      uint
	hards           [][]strategyTableAction
	splits          [][]strategyTableAction
	softs           [][]strategyTableAction
}

func newStratAction(dim1 int, dim2 int) [][]strategyTableAction {
	r := make([][]strategyTableAction, dim1)
	for i:=0;i<dim1;i++ {
		r[i] = make([]strategyTableAction, dim2)
	}
	return r
}

func NewDiscoveredStrategy(rules Rules, shoeFactory ShoeFactory, dealerStrategy PlayStrategy, iterations uint) *DiscoveredStrategy {
	return &DiscoveredStrategy{
		rules: rules,
		shoeFactory: shoeFactory,
		dealerStrategy: dealerStrategy,
		bettingStrategy: NewConsistentBettingStrategy(1),
		iterations: iterations,
		// 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
		hards: newStratAction(17, 11),
		// 2 3 4 5 6 7 8 9 10 A
		splits: newStratAction(10, 11),
		// 13 14 15 16 17 18 19 20 (A with any but another A or 10)
		softs: newStratAction(8, 11),
	}
}

func (this *DiscoveredStrategy) PrintStrategy(w io.Writer) {
	fmt.Fprintf(w, "%-15s", "")
	dealerUpCards := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for _, c := range dealerUpCards {
		fmt.Fprintf(w, "%-15s", Values()[c])
	}
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "Hard strategy table\n")

	for playerHandValue, perScoreHards := range this.hards {
		fmt.Fprintf(w, "%-15d", playerHandValue + 4)
		for _, dealerUpCardScore := range dealerUpCards {
			s := fmt.Sprintf("%s", perScoreHards[Values()[dealerUpCardScore].Index()])
			fmt.Fprintf(w, "%-15s", s)
		}
		fmt.Fprintf(w, "\n")
	}

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "Soft strategy table\n")
	for playerHandValue, perScoreSofts := range this.softs {
		fmt.Fprintf(w, "%-15d", playerHandValue + 13)
		for _, dealerUpCardScore := range dealerUpCards {
			s := fmt.Sprintf("%s", perScoreSofts[Values()[dealerUpCardScore].Index()])
			fmt.Fprintf(w, "%-15s", s)
		}
		fmt.Fprintf(w, "\n")
	}

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Splits strategy table\n")
	for i:=0;i<len(this.splits);i++ {
//	for playerSplitValue, perScoreSplits := range this.splits {
		perScoreSplits := this.splits[(i+1)%len(this.splits)]
		s := fmt.Sprintf("%c/%c", Values()[(i+1)%len(this.splits)].Symbol(), Values()[(i+1)%len(this.splits)].Symbol())
		fmt.Fprintf(w, "%-15s", s)
		for _, dealerUpCardScore := range dealerUpCards {
			s := fmt.Sprintf("%s", perScoreSplits[Values()[dealerUpCardScore].Index()])
			fmt.Fprintf(w, "%-15s", s)
		}
		fmt.Fprintf(w, "\n")
	}
}

func (this *DiscoveredStrategy) SetStrategy(currentHand Hand, dealerUpCard Card, gameAction GameAction) {
	if currentHand.Bust() {
		panic("You can't set a strategy for a bust hand!!!!")
	}
	if len(currentHand.Cards()) < 2 {
		panic("You can't set a strategy for a hand with less than two cards!!!")
	}
	if this.rules.CanSplit(currentHand) {
		this.splits[currentHand.FirstCard().Score() - 1][dealerUpCard.Score() - 1] = combineTableAction(this.splits[currentHand.FirstCard().Score() - 1][dealerUpCard.Score() - 1], gameAction)
	} else if currentHand.IsSoft() {
		// Smallest soft hand is 13
		this.softs[currentHand.Score() - 13][dealerUpCard.Score() - 1] = combineTableAction(this.softs[currentHand.Score() - 13][dealerUpCard.Score() - 1], gameAction)
	} else {
		// Smallest non split hard hand is 4 (Maybe you got 2/2 and split into 2/2 so many times you can't split again???)
		this.hards[currentHand.Score() - 4][dealerUpCard.Score() - 1] = combineTableAction(this.hards[currentHand.Score() - 4][dealerUpCard.Score() - 1], gameAction)
	}
}

func actionsClone(acts [][]strategyTableAction) [][]strategyTableAction {
	r := make([][]strategyTableAction, len(acts))
	for i:=0;i<len(acts);i++ {
		r[i] = make([]strategyTableAction, len(acts[i]))
		for j:=0;j<len(acts[i]);j++{
			r[i][j] = acts[i][j]
		}
	}
	return r
}

func (this *DiscoveredStrategy) Clone() *DiscoveredStrategy {
	return &DiscoveredStrategy{
		rules: this.rules,
		shoeFactory: this.shoeFactory,
		dealerStrategy: this.dealerStrategy,
		bettingStrategy: this.bettingStrategy,
		iterations: this.iterations,
		hards: actionsClone(this.hards),
		splits: actionsClone(this.splits),
		softs: actionsClone(this.softs),
	}
}



func (this *DiscoveredStrategy) TakeAction(currentHand Hand, dealerUpCard Card) GameAction {
	currentAction := this.NonRecursiveTakeAction(currentHand, dealerUpCard)
	if currentAction != nil {
		return currentAction
	}
	learnedAction, err := this.LearnAction(currentHand, dealerUpCard)
	if err != nil {
		panic(fmt.Sprintf("Logic error.  We shouldn't get errors here of %s", err))
	}
	if !this.rules.AllowsAction(learnedAction, currentHand) {
		panic(fmt.Sprintf("I am telling you to take an action you're not allowed to take %s with %s!", learnedAction, currentHand))
	}
	return learnedAction
}

func (this *DiscoveredStrategy) NonRecursiveTakeAction(currentHand Hand, dealerUpCard Card) GameAction {
	if currentHand.Bust() {
		panic("There is no action on a bust hand!")
	}
	if currentHand.Score() == 21 {
		// no silly moves on 21 even considered
		return STAND
	}
	var tableAction strategyTableAction
	if this.rules.CanSplit(currentHand) {
		tableAction = this.splits[currentHand.FirstCard().Score() - 1][dealerUpCard.Score() - 1]
	} else if currentHand.IsSoft() {
		tableAction = this.softs[currentHand.Score() - 13][dealerUpCard.Score() - 1]
	} else {
//		fmt.Printf("%d %d %d %d\n", currentHand.Score() - 4, dealerUpCard.Score() - 1, len(this.hards), len(this.hards[currentHand.Score() - 4]))
		tableAction = this.hards[currentHand.Score() - 4][dealerUpCard.Score() - 1]
	}
	if tableAction == nil {
		return nil
	} else {
		return tableAction.GameAction(this.rules, currentHand)
	}
}

func (this *DiscoveredStrategy) LearnAction(currentHand Hand, dealerUpCard Card) (GameAction, error) {
	var err error
	splitWinning := float64(math.MinInt64)
	doubleWinning := float64(math.MinInt64)
	hitWinning := float64(math.MinInt64)
	surrenderWinning := float64(math.MinInt64)
	standWinning := float64(math.MinInt64)
	var discoveredSplitTable *DiscoveredStrategy
	var discoveredDoubleTable *DiscoveredStrategy
	var discoveredSurrenderTable *DiscoveredStrategy
	var discoveredHitTable *DiscoveredStrategy
	var discoveredStandTable *DiscoveredStrategy

	currentAction := this.NonRecursiveTakeAction(currentHand, dealerUpCard)
	if currentAction != nil {
		return currentAction, nil
	}
	handDealer := NewForceDealerPlayerHands(currentHand, dealerUpCard.BlackjackValue())

	if this.rules.CanSplit(currentHand) {
		discoveredSplitTable = this.Clone()
		discoveredSplitTable.SetStrategy(currentHand, dealerUpCard, SPLIT)
		splitWinning, err = SimulateSingleHand(this.shoeFactory, handDealer, this.dealerStrategy, discoveredSplitTable, this.bettingStrategy, this.iterations, this.rules)
		if err != nil {
			return nil, err
		}
	}
	if this.rules.CanDouble(currentHand) {
		discoveredDoubleTable = this.Clone()
		discoveredDoubleTable.SetStrategy(currentHand, dealerUpCard, DOUBLE)
		doubleWinning, err = SimulateSingleHand(this.shoeFactory, handDealer, this.dealerStrategy, discoveredDoubleTable, this.bettingStrategy, this.iterations, this.rules)
		if err != nil {
			return nil, err
		}
	}
	if this.rules.CanSurrender(currentHand) {
		discoveredSurrenderTable = this.Clone()
		discoveredSurrenderTable.SetStrategy(currentHand, dealerUpCard, SURRENDER)
		surrenderWinning, err = SimulateSingleHand(this.shoeFactory, handDealer, this.dealerStrategy, discoveredSurrenderTable, this.bettingStrategy, this.iterations, this.rules)
		if err != nil {
			return nil, err
		}
	}
	discoveredHitTable = this.Clone()
	discoveredHitTable.SetStrategy(currentHand, dealerUpCard, HIT)
	hitWinning, err = SimulateSingleHand(this.shoeFactory, handDealer, this.dealerStrategy, discoveredHitTable, this.bettingStrategy, this.iterations, this.rules)
	if err != nil {
		return nil, err
	}
	discoveredStandTable = this.Clone()
	discoveredStandTable.SetStrategy(currentHand, dealerUpCard, STAND)
	standWinning, err = SimulateSingleHand(this.shoeFactory, handDealer, this.dealerStrategy, discoveredStandTable, this.bettingStrategy, this.iterations, this.rules)
	if err != nil {
		return nil, err
	}

	bestStrategy := STAND
	bestScore := standWinning
	if bestScore < hitWinning {
		bestStrategy = HIT
		bestScore = hitWinning
		this.hards = discoveredHitTable.hards
		this.softs = discoveredHitTable.softs
		this.splits = discoveredHitTable.splits
	}
	if bestScore < surrenderWinning {
		bestStrategy = SURRENDER
		bestScore = surrenderWinning
		this.hards = discoveredSurrenderTable.hards
		this.softs = discoveredSurrenderTable.softs
		this.splits = discoveredSurrenderTable.splits
	}
	if bestScore < doubleWinning {
		bestStrategy = DOUBLE
		bestScore = doubleWinning
		this.hards = discoveredDoubleTable.hards
		this.softs = discoveredDoubleTable.softs
		this.splits = discoveredDoubleTable.splits
	}
	if bestScore < splitWinning {
		bestStrategy = SPLIT
		bestScore = splitWinning
		this.hards = discoveredSplitTable.hards
		this.softs = discoveredSplitTable.softs
		this.splits = discoveredSplitTable.splits
	}
	//log.Printf("Best for %s vs %s is %s\n", currentHand, dealerUpCard, bestStrategy)
	if this.NonRecursiveTakeAction(currentHand, dealerUpCard) != bestStrategy {
		this.SetStrategy(currentHand, dealerUpCard, bestStrategy)
	}
	return bestStrategy, nil
}

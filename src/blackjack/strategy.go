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
}

type gameActionImpl struct {
	name   string
	symbol rune
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

var HIT = &gameActionImpl{"hit", 'H'}
var STAND = &gameActionImpl{"stand", 'S'}
var DOUBLE = &gameActionImpl{"double", 'D'}
var SPLIT = &gameActionImpl{"split", 'P'}
var SURRENDER = &gameActionImpl{"surrender", 'U'}

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
	hards           [][][]GameAction
	splits          [][][]GameAction
	softs           [][][]GameAction
}

func NewDiscoveredStrategy(rules Rules, shoeFactory ShoeFactory, dealerStrategy PlayStrategy, iterations uint) *DiscoveredStrategy {
	resets := make([][][][]GameAction, 3)
	for i := 0; i < len(resets); i++ {
		resets[i] = make([][][]GameAction, 21)
		for scores := 0; scores < len(resets[i]); scores++ {
			resets[i][scores] = make([][]GameAction, 11)
			for k :=0;k<len(resets[i][scores]);k++ {
				resets[i][scores][k] = []GameAction{}
			}
		}
	}

	return &DiscoveredStrategy{
		rules: rules,
		shoeFactory: shoeFactory,
		dealerStrategy: dealerStrategy,
		bettingStrategy: NewConsistentBettingStrategy(1),
		iterations: iterations,
		hards: resets[0],
		splits: resets[1],
		softs: resets[2],
	}
}

func (this *DiscoveredStrategy) PrintStrategy(w io.Writer) {
	for dealerUpCardScore, _ := range this.hards[0] {
		if dealerUpCardScore == 0 {
			fmt.Fprintf(w, "%-15s", "")
			continue
		}
		fmt.Fprintf(w, "%-15d", dealerUpCardScore)
	}

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Hard strategy table\n")
	for playerHandValue, perScoreHards := range this.hards {
		if playerHandValue < 5 {
			// Lowest hard hand is 5 because we consider 4 (2/2) a split hand
			continue
		}
		for dealerUpCardScore, perDealerCard := range perScoreHards {
			if dealerUpCardScore == 0 {
				fmt.Fprintf(w, "%-15d", playerHandValue)
				continue
			}
			s := fmt.Sprintf("%s", perDealerCard)
			fmt.Fprintf(w, "%-15s", s)
		}
		fmt.Fprintf(w, "\n")
	}

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Soft strategy table\n")
	for playerHandValue, perScoreHards := range this.softs {
		if playerHandValue < 13 {
			// Lowest soft hand is 13 because we consider 12 (A/A) a split hand)
			continue
		}
		for dealerUpCardScore, perDealerCard := range perScoreHards {
			if dealerUpCardScore == 0 {
				fmt.Fprintf(w, "%-15d", playerHandValue)
				continue
			}
			s := fmt.Sprintf("%s", perDealerCard)
			fmt.Fprintf(w, "%-15s", s)
		}
		fmt.Fprintf(w, "\n")
	}

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Splits strategy table\n")
	for playerSplitValue, perScoreHards := range this.splits {
		if playerSplitValue > 11 || playerSplitValue == 0 {
			continue
		}
		for dealerUpCardScore, perDealerCard := range perScoreHards {
			if dealerUpCardScore == 0 {
				s := fmt.Sprintf("%d/%d", playerSplitValue,playerSplitValue)
				fmt.Fprintf(w, "%-15s", s)
				continue
			}
			s := fmt.Sprintf("%s", perDealerCard)
			fmt.Fprintf(w, "%-15s", s)
		}
		fmt.Fprintf(w, "\n")
	}
}

func (this *DiscoveredStrategy) SetStrategy(currentHand Hand, dealerUpCard Card, gameAction GameAction) {
	if this.rules.CanSplit(currentHand) {
		for _, v := range this.splits[currentHand.FirstCard().Score()][dealerUpCard.Score()] {
			if v == gameAction {
				// Already set
				return
			}
		}
		this.splits[currentHand.FirstCard().Score()][dealerUpCard.Score()] = append(this.splits[currentHand.FirstCard().Score()][dealerUpCard.Score()], gameAction)
	} else if currentHand.IsSoft() {
		for _, v := range this.softs[currentHand.Score()][dealerUpCard.Score()] {
			if v == gameAction {
				// Already set
				return
			}
		}
		this.softs[currentHand.Score()][dealerUpCard.Score()] = append(this.softs[currentHand.Score()][dealerUpCard.Score()], gameAction)
	} else {
		for _, v := range this.hards[currentHand.Score()][dealerUpCard.Score()] {
			if v == gameAction {
				// Already set
				return
			}
		}
		this.hards[currentHand.Score()][dealerUpCard.Score()] = append(this.hards[currentHand.Score()][dealerUpCard.Score()], gameAction)
	}
}

func (this *DiscoveredStrategy) Clone() *DiscoveredStrategy {
	thisResets := [][][][]GameAction{this.hards, this.splits, this.softs}
	resets := make([][][][]GameAction, 3)
	for i := 0; i < len(resets); i++ {
		resets[i] = make([][][]GameAction, 21)
		for scores := 0; scores < len(resets[i]); scores++ {
			resets[i][scores] = make([][]GameAction, 11)
			for j := 0; j < len(resets[i][scores]); j++ {
				resets[i][scores][j] = make([]GameAction, len(thisResets[i][scores][j]))
				for k :=0;k<len(resets[i][scores][j]);k++ {
					resets[i][scores][j][k] = thisResets[i][scores][j][k]
				}
			}
		}
	}

	return &DiscoveredStrategy{
		rules: this.rules,
		shoeFactory: this.shoeFactory,
		dealerStrategy: this.dealerStrategy,
		bettingStrategy: this.bettingStrategy,
		iterations: this.iterations,
		hards: resets[0],
		splits: resets[1],
		softs: resets[2],
	}
}



func (this *DiscoveredStrategy) TakeAction(currentHand Hand, dealerUpCard Card) GameAction {
	if currentHand.Score() >= 21 {
		return STAND
	}
	var actions []GameAction
	if this.rules.CanSplit(currentHand) {
		actions = this.splits[currentHand.FirstCard().Score()][dealerUpCard.Score()]
	} else if currentHand.IsSoft() {
		actions = this.softs[currentHand.Score()][dealerUpCard.Score()]
	} else {
		actions = this.hards[currentHand.Score()][dealerUpCard.Score()]
	}
	for _, action := range actions {
		if this.rules.AllowsAction(action, currentHand) {
			return action
		}
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

func (this *DiscoveredStrategy) nonRecursiveTakeAction(currentHand Hand, dealerUpCard Card) GameAction {
	if currentHand.Score() >= 21 {
		return STAND
	}
	var actions []GameAction
	if this.rules.CanSplit(currentHand) {
		actions = this.splits[currentHand.FirstCard().Score()][dealerUpCard.Score()]
	} else if currentHand.IsSoft() {
		actions = this.softs[currentHand.Score()][dealerUpCard.Score()]
	} else {
		actions = this.hards[currentHand.Score()][dealerUpCard.Score()]
	}
	for _, action := range actions {
		if this.rules.AllowsAction(action, currentHand) {
			return action
		}
	}
	return nil
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

	currentAction := this.nonRecursiveTakeAction(currentHand, dealerUpCard)
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
	this.SetStrategy(currentHand, dealerUpCard, bestStrategy)
	return bestStrategy, nil
}

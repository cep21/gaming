/**
 * Date: 4/1/14
 * Time: 3:46 PM
 * @author jack 
 */
package blackjack

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

func (this* dealerHitStrategy) TakeAction(currentHand Hand, _ Card) GameAction {
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

func (this* alwaysHitStrategy) TakeAction(_ Hand, _ Card) GameAction {
	return HIT;
}

func NewAlwaysHitStrategy() PlayStrategy {
	return &alwaysHitStrategy{}
}

// Always stand
type alwaysStandStrategy struct {
}

func (this* alwaysStandStrategy) TakeAction(_ Hand, _ Card) GameAction {
	return STAND;
}

func NewAlwaysStandStrategy() PlayStrategy {
	return &alwaysStandStrategy{}
}

// Never bust
type neverBustStrategy struct {
	shouldHitSoft bool
}

func (this* neverBustStrategy) TakeAction(currentHand Hand, _ Card) GameAction {
	if (currentHand.IsSoft() && this.shouldHitSoft) || currentHand.Score() < 12 {
		return HIT;
	} else {
		return STAND;
	}
}

func NewNeverBustStrategy(should_hit_soft bool) PlayStrategy {
	return &neverBustStrategy{should_hit_soft}
}

func PlayHandOnStrategy(currentHand Hand, shownCard Card, strategy PlayStrategy, deck Shoe)  {
	for ; deck.CardsLeft() != 0; {
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

func (this* hitOnAScoreStrategy) TakeAction(currentHand Hand, _ Card) GameAction {
	if (currentHand.IsSoft() && currentHand.Score() <= this.softScoreToHit) || (!currentHand.IsSoft() && currentHand.Score() <= this.hardScoreToHit) {
		return HIT
	} else {
		return STAND
	}
}

func NewHitOnAScoreStrategy(soft_score_to_hit uint, hard_score_to_hit uint) PlayStrategy {
	return &hitOnAScoreStrategy{hard_score_to_hit, soft_score_to_hit}
}

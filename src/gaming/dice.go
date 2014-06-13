package gaming

import (
	"math/rand"
	"time"
)

type Dice interface {
	Sides() uint
	Roll() uint
}

func NewDice(sides uint) Dice {
	return NewDiceSeed(sides, time.Now().UnixNano())
}

func NewDiceSeed(sides uint, seed int64) Dice {
	return &diceImpl{sides: sides, rand: rand.New(rand.NewSource(seed))}
}



/// --- private implementations

type diceImpl struct {
	sides uint
	rand  *rand.Rand
}

func (this *diceImpl) Roll() uint {
	return uint(this.rand.Intn(int(this.sides)))
}

func (this *diceImpl) Sides() uint {
	return this.sides
}


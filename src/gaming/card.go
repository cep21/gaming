package gaming

type Suit interface {
	Name() string
	Symbol() rune
	Index() uint
}

type suitImpl struct {
	name   string
	symbol rune
	index  uint
}

func (this *suitImpl) Name() string {
	return this.name
}

func (this *suitImpl) Symbol() rune {
	return this.symbol
}

func (this *suitImpl) Index() uint {
	return this.index
}

var Spade = &suitImpl{name: "spade", symbol: 's', index: 0}
var Club = &suitImpl{name: "club", symbol: 'c', index: 1}
var Heart = &suitImpl{name: "heart", symbol: 'h', index: 2}
var Diamond = &suitImpl{name: "diamond", symbol: 'd', index: 3}

func Suits() []Suit {
	return []Suit{Spade, Club, Heart, Diamond}
}

type Value interface {
	Name() string
	Symbol() rune
	Index() uint
}

type valueImpl struct {
	name   string
	symbol rune
	index  uint
}

func (this *valueImpl) Name() string {
	return this.name
}

func (this *valueImpl) Symbol() rune {
	return this.symbol
}

func (this *valueImpl) Index() uint {
	return this.index
}

var Ace = &valueImpl{name:"ace", symbol:'a', index: 0}
var Two = &valueImpl{name:"ace", symbol:'2', index: 1}
var Three = &valueImpl{name:"ace", symbol:'3', index: 2}
var Four = &valueImpl{name:"ace", symbol:'4', index: 3}
var Five = &valueImpl{name:"ace", symbol:'5', index: 4}
var Six = &valueImpl{name:"ace", symbol:'6', index: 5}
var Seven = &valueImpl{name:"ace", symbol:'7', index: 6}
var Eight = &valueImpl{name:"ace", symbol:'8', index: 7}
var Nine = &valueImpl{name:"ace", symbol:'9', index: 8}
var Ten = &valueImpl{name:"ace", symbol:'t', index: 9}
var Jack = &valueImpl{name:"ace", symbol:'j', index: 10}
var Queen = &valueImpl{name:"ace", symbol:'q', index: 11}
var King = &valueImpl{name:"ace", symbol:'k', index: 12}
var Joker = &valueImpl{name:"joker", symbol:'o', index: 13}

func Values() []Value {
	return []Value{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
}

func JokerValues() []Value {
	return []Value{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Joker}
}

type Card interface {
	Suit() Suit
	Value() Value
}

type cardImpl struct {
	suit    Suit
	value   Value
}

func (this *cardImpl) Suit() Suit {
	return this.suit
}

func (this *cardImpl) Value() Value {
	return this.value
}

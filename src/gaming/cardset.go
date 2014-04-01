package gaming

type CardSet interface {
	Cards() []Card
	Peek() Card
	Pop() Card
	CardsLeft() uint
	IsEmpty() bool
}

type cardSetImpl struct {
	cards []Card
}

func (this *cardSetImpl) Cards() []Card {
	return this.cards
}

func (this *cardSetImpl) Peek() Card {
	return this.cards[0]
}

func (this *cardSetImpl) Pop() Card {
	c := this.cards[0]
	this.cards = this.cards[1:]
	return c
}

func (this *cardSetImpl) CardsLeft() uint {
	return uint(len(this.cards))
}

func (this *cardSetImpl) IsEmpty() bool {
	return this.CardsLeft() != 0
}

func NewDeck(cards []Card) CardSet {
	return &cardSetImpl{cards}
}

func SingleDeck() CardSet {
	cards := []Card{}
	for _, suit := range Suits() {
		for _, value := range Values() {
			cards = append(cards, &cardImpl{suit: suit, value: value})
		}
	}

	return NewDeck(cards)
}

func JokerDeck(num_jokers int) CardSet {
	cards := SingleDeck().Cards()
	for i := 0; i < num_jokers; i++ {
		cards = append(cards, &cardImpl{suit: nil, value: Joker})
	}
	return NewDeck(cards)
}

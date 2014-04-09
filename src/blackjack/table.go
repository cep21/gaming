package blackjack

import "gaming/bankroll"

type Table interface {
	Spots() uint
	Player(spot uint) Player
	SetPlayer(player Player, spot uint) Player
	RemovePlayer(spot uint) Player
	PlayRound() error
	Dealer() Dealer
	ShoeFactory() ShoeFactory
	activePlayers() []Player
}

type tableImpl struct {
	spots       uint
	players     []Player
	dealer      Dealer
	shoeFactory ShoeFactory
	rules       Rules
}

func NewTable(dealer Dealer, shoeFactory ShoeFactory, spots uint, rules Rules) Table {
	return &tableImpl{
		spots:       spots,
		players:     make([]Player, spots),
		dealer:      dealer,
		shoeFactory: shoeFactory,
		rules:       rules,
	}
}

func (table *tableImpl) Spots() uint {
	return table.spots
}
func (table *tableImpl) Player(spot uint) Player {
	return table.players[spot]
}
func (table *tableImpl) activePlayers() []Player {
	ret := []Player{}
	for _, p := range table.players {
		if p != nil {
			ret = append(ret, p)
		}
	}
	return ret
}
func (table *tableImpl) SetPlayer(player Player, spot uint) Player {
	oldPlayer := table.players[spot]
	table.players[spot] = player
	return oldPlayer
}
func (table *tableImpl) RemovePlayer(spot uint) Player {
	return table.SetPlayer(nil, spot)
}

func assertHandBankrollsAreEmpty(hands []Hand) {
	for _, h := range hands {
		if h.MoneyInThisHand().CurrentBankroll() != 0 {
			panic("Logic error: Hands should be empty at the end")
		}
	}
}

func (table *tableImpl) PlayRound() error {
	shoe := table.shoeFactory.CreateShoe()
	activePlayers := table.activePlayers()
	err := table.dealer.HandDealer().DealHands(shoe, activePlayers, table.dealer)
	if err != nil {
		return err
	}
	defer assertHandBankrollsAreEmpty(table.Dealer().Hands())
	dealerHand := table.Dealer().Hands()[0]

	if dealerHand.IsBlackjack() {
		// TODO(insurance)
	}
	for _, activePlayer := range activePlayers {
		defer assertHandBankrollsAreEmpty(activePlayer.Hands())
		survivingPlayerHands := []Hand{}
		for _, playerHand := range activePlayer.Hands() {
			if playerHand.IsBlackjack() {
				if dealerHand.IsBlackjack() {
					// Push the money back
					playerHand.MoneyInThisHand().TransferMoneyTo(activePlayer.Bankroll(), playerHand.MoneyInThisHand().CurrentBankroll())
				} else {
					// Player wins: pay bonus
					table.dealer.Bankroll().TransferMoneyTo(activePlayer.Bankroll(), playerHand.MoneyInThisHand().CurrentBankroll()*bankroll.Money(table.rules.BlackjackPayout()))
					// Give me my money back!
					playerHand.MoneyInThisHand().TransferMoneyTo(activePlayer.Bankroll(), playerHand.MoneyInThisHand().CurrentBankroll())
				}
			} else if dealerHand.IsBlackjack() {
				// Dealer wins
				playerHand.MoneyInThisHand().TransferMoneyTo(table.dealer.Bankroll(), playerHand.MoneyInThisHand().CurrentBankroll())
			} else {
				survivingPlayerHands = append(survivingPlayerHands, playerHand)
			}
		}
		activePlayer.SetHands(survivingPlayerHands)
	}

	for _, activePlayer := range activePlayers {
		finalPlayerHands := []Hand{}
		for _, playerHand := range activePlayer.Hands() {
			thisHandsNewHands, err := PlayHand(playerHand, dealerHand, activePlayer, table.rules, table.dealer.Bankroll(), shoe)
			if err != nil {
				return err
			}
			defer assertHandBankrollsAreEmpty(thisHandsNewHands)
			for _, hand := range thisHandsNewHands {
				if hand.Bust() {
					hand.MoneyInThisHand().TransferMoneyTo(table.dealer.Bankroll(), hand.MoneyInThisHand().CurrentBankroll())
				} else {
					finalPlayerHands = append(finalPlayerHands, hand)
				}
			}
		}
		activePlayer.SetHands(finalPlayerHands)
	}

	allDealerHands, err := PlayHand(dealerHand, nil, table.dealer, table.rules, table.dealer.Bankroll(), shoe)
	if err != nil {
		return err
	}
	defer assertHandBankrollsAreEmpty(allDealerHands)
	if len(allDealerHands) != 1 || allDealerHands[0] != dealerHand {
		return ERR_INVALID_DEALER_OPERATION
	}

	for _, activePlayer := range activePlayers {
		for _, finalPlayerHand := range activePlayer.Hands() {
			if dealerHand.Bust() {
				if finalPlayerHand.Bust() {
					panic("Logic error: This should never happen")
				}
				table.dealer.Bankroll().TransferMoneyTo(activePlayer.Bankroll(), finalPlayerHand.MoneyInThisHand().CurrentBankroll())
				finalPlayerHand.MoneyInThisHand().TransferMoneyTo(activePlayer.Bankroll(), finalPlayerHand.MoneyInThisHand().CurrentBankroll())
			} else if dealerHand.Score() > finalPlayerHand.Score() {
				// You loose, give it to the house
				finalPlayerHand.MoneyInThisHand().TransferMoneyTo(table.dealer.Bankroll(), finalPlayerHand.MoneyInThisHand().CurrentBankroll())
			} else if dealerHand.Score() < finalPlayerHand.Score() {
				// You win.  Take some from the house
				table.dealer.Bankroll().TransferMoneyTo(activePlayer.Bankroll(), finalPlayerHand.MoneyInThisHand().CurrentBankroll())
				finalPlayerHand.MoneyInThisHand().TransferMoneyTo(activePlayer.Bankroll(), finalPlayerHand.MoneyInThisHand().CurrentBankroll())
			} else {
				// Push the money back
				finalPlayerHand.MoneyInThisHand().TransferMoneyTo(activePlayer.Bankroll(), finalPlayerHand.MoneyInThisHand().CurrentBankroll())
			}
		}
	}
	return nil
}
func (table *tableImpl) Dealer() Dealer {
	return table.dealer
}
func (table *tableImpl) ShoeFactory() ShoeFactory {
	return table.shoeFactory
}

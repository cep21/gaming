/**
 * Date: 4/6/14
 * Time: 10:51 AM
 * @author jack
 */
package blackjack

//import "errors"

//var errUnableToSplit = errors.New("Unable to split this hand")
//
//type PlayedHand interface {
//	Hand() Hand
//	MoneyInThisHand() MoneyHolder
//	SplitNumber() uint
//	LastAction() GameAction
//	SetLastAction(GameAction)
//	SplitHand(bankrollToDrawFrom MoneyHolder) (PlayedHand, PlayedHand, error)
//}
//
//type playedHandImpl struct {
//	hand            Hand
//	moneyInThisHand MoneyHolder
//	splitNumber     uint
//	lastAction      GameAction
//}
//
//func (this *playedHandImpl) Hand() Hand {
//	return this.hand
//}
//
//func (this *playedHandImpl) MoneyInThisHand() MoneyHolder {
//	return this.moneyInThisHand
//}
//
//func (this *playedHandImpl) SplitNumber() uint {
//	return this.splitNumber
//}
//
//func (this *playedHandImpl) LastAction() GameAction {
//	return this.lastAction
//}
//
//func (this *playedHandImpl) SplitHand(bankrollToDrawFrom MoneyHolder) (PlayedHand, PlayedHand, error) {
//	if !this.Hand().CanSplit() {
//		return nil, nil, errUnableToSplit
//	}
//	cardsInHand := this.Hand().Cards()
//	secondHandMoney := NewMoneyHolder()
//	bankrollToDrawFrom.TransferMoneyTo(secondHandMoney, this.MoneyInThisHand().CurrentBankroll())
//	firstHand := NewHand(cardsInHand[0])
//	secondHand := NewHand(cardsInHand[1])
//	return NewPlayedHand(firstHand, this.moneyInThisHand, this.splitNumber + 1), NewPlayedHand(secondHand, secondHandMoney, this.splitNumber + 1), nil
//}
//
//func NewPlayedHand(hand Hand, moneyInThisHand MoneyHolder, splitNumber uint) PlayedHand {
//	return &playedHandImpl{
//		hand: hand,
//		moneyInThisHand: moneyInThisHand,
//		splitNumber: splitNumber,
//	}
//}

/**
 * Date: 4/1/14
 * Time: 6:07 PM
 * @author jack 
 */
package blackjack

type GameTable interface {
	CardSet() CardSet
	NumPlayers() uint // Not counting the dealer
	DealerHitStrategy() ShouldHitStrategy
	PlayerHitStrategy(player_number uint) ShouldHitStrategy
}

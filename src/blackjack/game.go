/**
 * Date: 4/1/14
 * Time: 5:48 PM
 * @author jack 
 */
package blackjack

type PlayGame interface {
	CardSet() CardSet
	NumPlayers() uint // Not counting the dealer
	DealerHitStrategy() ShouldHitStrategy
	PlayerHitStrategy(player_number uint) ShouldHitStrategy
}

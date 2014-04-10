/**
 * Date: 4/9/14
 * Time: 2:16 PM
 * @author jack
 */
package main

import (
	"gaming"
	"blackjack"
	"flag"
	"os"
	"math/rand"
	"time"
)

func main() {
	surrenderAllowed := flag.Bool("surrender", false, "Is late surrender allowed?")
	doubleAfterSplitAllowed := flag.Bool("doublesplit", false, "Is double after split allowed?")
	hitSoft17 := flag.Bool("soft17", false, "Does dealer hit soft 17?")
	iterations := flag.Uint("iterations", 20000, "Number of iterations to simulate")
	decks := flag.Uint("decks", 0, "Number of decks to use (0 means infinite)")
	seed := flag.Int("seed", 0, "Random seed for simulation")
	flag.Parse()

	rulesFactory := blackjack.NewRulesetFactory()
	rulesFactory.DealerHitsSoft17(*hitSoft17)
	rulesFactory.DoubleAfterSplit(*doubleAfterSplitAllowed)
	if *surrenderAllowed {
		rulesFactory.SurrenderOption(blackjack.LATE_SURRENDER)
	}
	rules := rulesFactory.Build()
	if *seed == 0 {
		*seed = int(time.Now().Unix())
	}
	r := rand.New(rand.NewSource(int64(*seed)))
	var shoeFactory blackjack.ShoeFactory
	if *decks == 0 {
		shoeFactory = blackjack.NewClonedDeckFactory(blackjack.NewInfiniteShoe(r), r)
	} else {
		shoeFactory = blackjack.NewClonedDeckFactory(blackjack.Decks(*decks), r)
	}

	dealerStrategy := blackjack.NewDealerStrategy(rules.DealerHitOnSoft17())
	strat := blackjack.NewDiscoveredStrategy(rules, shoeFactory, dealerStrategy, *iterations)
	for isSplit := 0 ;isSplit < 2; isSplit++ {
		for isSoft := 0; isSoft < 2; isSoft ++ {
			for totalValue := uint(21) ; totalValue > uint(0); totalValue-- {
				for _, cardOne := range blackjack.Values() {
					for _, cardTwo := range blackjack.Values() {
						for _, dealerUpCard := range blackjack.Values() {
							playerHand := blackjack.NewHand(blackjack.NewCard(gaming.Spade, cardOne), blackjack.NewCard(gaming.Spade, cardTwo))
							dealerCard := blackjack.NewCard(gaming.Heart, dealerUpCard)
							if playerHand.Score() == totalValue && playerHand.IsSoft() == (isSoft == 1) && rules.CanSplit(playerHand) == (isSplit == 1) {
								// Learn higher down to lower
								strat.TakeAction(playerHand, dealerCard)
								strat.PrintStrategy(os.Stdout)
							}
						}
					}
				}
			}
		}
	}
}

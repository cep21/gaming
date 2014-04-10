/**
 * Date: 4/9/14
 * Time: 2:16 PM
 * @author jack 
 */
package main

import "fmt"
import "blackjack"

func main() {
	numberOfIterations := 10000
	rules := blackjack.NewRulesetFactory().Build()
	strategy := NewDiscoveredStrategy(rules)
}



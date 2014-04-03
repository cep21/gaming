/**
 * Date: 4/3/14
 * Time: 1:28 PM
 * @author jack 
 */
package blackjack

import (
	"testing"
)

func TestBettingConstant(t *testing.T) {
	strat := NewConsistentBettingStrategy(2)
	if strat.GetUnitsToBet() != 2 {
		t.Error("Expected 2 units to bet")
	}
}


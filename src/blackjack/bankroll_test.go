/**
 * Date: 4/4/14
 * Time: 10:57 AM
 * @author jack 
 */
package blackjack

import "testing"

func TestBankrollChange(t *testing.T) {
	b := NewMoneyHolder()
	if b.CurrentBankroll() != 0 {
		t.Error("I expected a bankroll of 0!")
	}
}


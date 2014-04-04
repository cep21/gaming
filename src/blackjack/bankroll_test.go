/**
 * Date: 4/4/14
 * Time: 10:57 AM
 * @author jack 
 */
package blackjack

import "testing"

func TestBankrollChange(t *testing.T) {
	b := NewBankroll(10)
	b.ChangeBankroll(1)
	if b.CurrentBankroll() != 11 {
		t.Error("I expected a bankroll of 11!")
	}
}


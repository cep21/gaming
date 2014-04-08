/**
 * Date: 4/4/14
 * Time: 10:57 AM
 * @author jack 
 */
package bankroll

import "testing"

func TestBankrollChange(t *testing.T) {
	b := NewMoneyHolder()
	if b.CurrentBankroll() != 0 {
		t.Error("I expected a bankroll of 0!")
	}
	d := NewMoneyHolder()
	b.TransferMoneyTo(d, 10)
	if b.CurrentBankroll() != -10 {
		t.Error("Expected a final bankroll of -10")
	}
	if d.CurrentBankroll() != 10 {
		t.Error("Expected a final bankroll of 10")
	}
}


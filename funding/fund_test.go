package funding

import (
	"fmt"
	"testing"
)

func BenchmarkFund(b *testing.B) {
	// Add as many dollars as we have iterations this run
	fund := NewFund(b.N)

	// Burn through them one at a time until they are all gone
	for i := 0; i < b.N; i++ {
		fund.Withdraw(1)
	}

	if fund.Balance() != 0 {
		b.Error("Balance wasn't zero:", fund.Balance())
	}
}

func TestFundWithdraw(t *testing.T) {
	initialBalance := 100
	fund := Fund{
		balance: initialBalance,
	}

	balance, err := fund.Withdraw(50)

	if err != nil {
		t.Fatalf("Withdraw failed: %s", err)
	}
	fmt.Printf("Withdraw test successful. Balance: %d - ", balance)
}

func TestFundBalance(t *testing.T) {
	initialBalance := 100
	fund := Fund{
		balance: initialBalance,
	}

	if fund.Balance() != 100 {
		t.Fatalf("Error fetching balance: %d", initialBalance)
	}
}

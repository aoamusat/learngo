package funding

import (
	"fmt"
	"sync"
	"testing"
)

const WORKERS = 10

func BenchmarkFund(b *testing.B) {
	// Skip N = 1
	if b.N < WORKERS {
		return
	}

	// Add as many dollars as we have iterations this run
	fund := NewFund(b.N)

	dollarsPerFounder := int(b.N / WORKERS)

	var wg sync.WaitGroup

	// Burn through them one at a time until they are all gone
	for i := 0; i < WORKERS; i++ {
		// Let the waitgroup know we're adding a goroutine
		wg.Add(1)
		// Spawn off a founder worker, as a closure
		go func() {
			// Mark this worker done when the function finishes
			defer wg.Done()

			for j := 0; j < dollarsPerFounder; j++ {
				fund.Withdraw(1)
			}

		}() // Remember to call the closure!
	}

	// Wait for all the workers to finish
	wg.Wait()

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

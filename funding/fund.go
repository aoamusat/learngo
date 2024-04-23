package funding

import "errors"

type Fund struct {
	// balance is unexported (private), because it's lowercase
	balance int
}

// A regular function returning a pointer to a fund
func NewFund(initialBalance int) *Fund {
	// We can return a pointer to a new struct without worrying about
	// whether it's on the stack or heap: Go figures that out for us.

	newFund := Fund{balance: initialBalance}
	var newFundPtr *Fund = &newFund

	return newFundPtr
}

// Defines Balance method on the Fund Type
func (f *Fund) Balance() int {
	return f.balance
}

// Defines Withdraw method on the Fund Type
func (f *Fund) Withdraw(amount int) (int, error) {
	accountBalance := f.Balance()
	if amount > accountBalance {
		return accountBalance, errors.New("insufficient balance")
	}

	f.balance -= amount

	return f.Balance(), nil
}

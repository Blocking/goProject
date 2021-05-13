package bankcore

import (
	"errors"
	"fmt"
)

func Hello() string {
	return "Hey! I'm working! "
}

// Customer ...
type Customer struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

// Account ...
type Account struct {
	Customer
	Number  int32   `json:"number"`
	Balance float64 `json:"balance"`
}

func (a *Account) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to withdraw should be greater than the account's balance")
	}

	a.Balance -= amount
	return nil
}

// Statement ...
func (a *Account) Statement() string {
	/*bytes, err := json.Marshal(a)
	if err != nil{
		panic(err)
	}
	return string(bytes)*/
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

func (a *Account) Transfer(to *Account, amount float64) error {
	if a.Balance < amount {
		return errors.New(a.Statement() + " money is not enough")
	}

	err := a.Withdraw(amount)
	if err != nil {
		return err
	}
	err = to.Deposit(amount)
	if err != nil {
		return err
	}
	return nil
}

// Bank ...
type Bank interface {
	Statement() string
}

// Statement ...
func Statement(b Bank) string {
	return b.Statement()
}

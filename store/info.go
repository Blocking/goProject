package store

import "fmt"

type Account struct {
	lastName  string
	firstName string
}

func (a *Account) changeName(newName string) {
	a.firstName = newName
}

func (a *Account) String() string {
	return fmt.Sprintf("%v  %v", a.firstName, a.lastName)
}

type Employee struct {
	Account
	Credits float64
}

func (e *Employee) AddCredits(size float64) {
	e.Credits += size
}
func (e *Employee) RemoveCredits() {
	e.Credits = 0
}
func (e *Employee) CheckCredits() {
	fmt.Println("amount:", e.Credits)
}

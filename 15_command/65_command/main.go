package main

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Printf("Deposited %d, balance is now %d\n", amount, b.balance)
}

func (b *BankAccount) Withdraw(amount int) {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		fmt.Printf("Withdrew %d, balance is now %d\n", amount, b.balance)
	}
}

// Command will handle itself
type Command interface {
	Call()
}

type Action int

const (
	Deposit  Action = 1
	Withdraw Action = 2
)

type BankAccountCommand struct {
	account *BankAccount
	action  Action
	amount  int
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account, action, amount}
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
	case Withdraw:
		b.account.Withdraw(b.amount)
	}
}

func main() {
	ba := BankAccount{}
	cmd := BankAccountCommand{&ba, Deposit, 100}
	cmd.Call()
	fmt.Println(ba)
	cmd2 := BankAccountCommand{&ba, Withdraw, 50}
	cmd2.Call()
	fmt.Println(ba)
}

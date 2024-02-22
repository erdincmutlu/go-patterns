package main

import "fmt"

type BankAccount struct {
	Balance int
}

func Deposit(ba *BankAccount, amount int) {
	fmt.Printf("Depositing %d\n", amount)
	ba.Balance += amount
}

func Withdraw(ba *BankAccount, amount int) {
	fmt.Printf("Withdrawing %d\n", amount)
	ba.Balance -= amount
}

func main() {
	ba := &BankAccount{0}
	var commands []func()
	commands = append(commands, func() {
		Deposit(ba, 100)
	})
	commands = append(commands, func() {
		Withdraw(ba, 25)
	})

	// Put into commands list so we can run all commands
	for _, cmd := range commands {
		cmd()
	}

	fmt.Println(*ba)
}

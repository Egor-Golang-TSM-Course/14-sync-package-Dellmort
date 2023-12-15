package task1

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	sync.Mutex
	balance int
}

func NewBankAccount() *BankAccount {
	return new(BankAccount)
}

func (b *BankAccount) Deposit(amount int) {
	b.Lock()
	defer b.Unlock()

	b.balance += amount
}

func (b *BankAccount) Withdraw(amount int) {
	b.Lock()
	defer b.Unlock()

	if b.balance-amount < 0 {
		fmt.Println("Произошла ошибка, недостаточно денег для снятия")
		return
	}

	b.balance -= amount
}

func (b *BankAccount) Balance() {
	b.Lock()
	defer b.Unlock()

	fmt.Println(b.balance)
}

func Start() {
	account := NewBankAccount()

	go account.Deposit(10000)
	go account.Withdraw(500)
	go account.Withdraw(333)
	go account.Withdraw(5200)

	time.Sleep(1 * time.Second)
	account.Balance() // 3967
}

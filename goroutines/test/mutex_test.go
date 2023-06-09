package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// solusi dari menghindari race condition
// walaupun secara kinerja lebih cepatan tanpa locking tapi ini lebih aman (ini juga tidak termasuk lambat, masih kategori cepat)
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("counter = ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	// untuk write locking
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	// untuk write unlock
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	// untuk read locking
	account.RWMutex.RLock()
	balance := account.Balance
	// untuk read unlock
	account.RWMutex.RUnlock()

	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("total balance = ", account.GetBalance())
}

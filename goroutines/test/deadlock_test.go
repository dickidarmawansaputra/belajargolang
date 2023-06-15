package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	Mutex   sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(userSatu *UserBalance, userDua *UserBalance, amount int) {
	userSatu.Lock()
	fmt.Println("Lock user 1", userSatu.Name)
	userSatu.Change(-amount)
	time.Sleep(1 * time.Second)

	userDua.Lock()
	fmt.Println("Lock user 2", userDua.Name)
	userDua.Change(amount)
	time.Sleep(1 * time.Second)

	userSatu.Unlock()
	userDua.Unlock()

}

func TestDeadlock(t *testing.T) {
	userSatu := UserBalance{
		Name:    "Dicki",
		Balance: 10000,
	}

	userDua := UserBalance{
		Name:    "Saputra",
		Balance: 20000,
	}

	// goroutine 1 & 2 saling menunggu. sehingga ada proses yg belum selesai
	// user 1 mau lock user 2 tapi
	go Transfer(&userSatu, &userDua, 10000)
	// user 2 sudah keburu di lock sama goroutine ini
	go Transfer(&userDua, &userSatu, 10000)

	time.Sleep(3 * time.Second)

	fmt.Println("user: ", userSatu.Name, "balance: ", userSatu.Balance)
	fmt.Println("user: ", userDua.Name, "balance: ", userDua.Balance)
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	deposits = make(chan int64)
	balances = make(chan int64)
	withdraw = make(chan int64)
	done     = make(chan bool)
)

func teller() {
	var balance int64
	for {
		select {
		case amount := <-withdraw:
			remnant := balance - amount
			if remnant > 0 {
				balance = remnant
				done <- true
			} else {
				done <- false
			}
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}
func Withdraw(amount int64) bool { withdraw <- amount; return <-done }
func Deposit(amount int64)       { deposits <- amount }
func Balance() int64             { return <-balances }

func main() {
	go teller()

	rand.Seed(time.Now().UnixNano())
	count := 100
	realAmount := int64(0)
	for i := 0; i < count; i++ {
		fund := rand.Int63n(100)
		realAmount += fund
		go Deposit(fund)
	}

	time.Sleep(1 * time.Second) // time for goroutines finish their work
	fmt.Println("Balance(): ", Balance())
	fmt.Println("realAmount: ", realAmount)
	fmt.Println(Withdraw(500))
	fmt.Println("Balance(): ", Balance())
}

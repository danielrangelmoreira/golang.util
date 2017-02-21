package main

import (
	"fmt"
	"math/rand"
	//"runtime"
	"sync"
	"time"
)

var balance int64

var mu sync.Mutex

func Deposit(amount int64) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int64 {
	return balance
}

func main() {
	rand.Seed(time.Now().UnixNano())
	count := 100000
	realAmount := int64(0)
	for i := 0; i < count; i++ {
		deposit := rand.Int63n(100)
		realAmount += deposit
		Deposit(deposit)
	}
	//time.Sleep(3 * time.Second)
	fmt.Println("Balance(): ", Balance())
	fmt.Println("realAmount: ", realAmount)
}

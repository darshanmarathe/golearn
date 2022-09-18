package main

import (
	"fmt"
	"sync"
)

var (
	balance int
	mutex   sync.Mutex
)

func deposit(amt int, wg *sync.WaitGroup) {
	mutex.Lock()
	balance += amt
	fmt.Println("Depositing amount ", amt, " and balance", balance)
	mutex.Unlock()
	wg.Done()
}

func withdraw(amt int, wg *sync.WaitGroup) {
	mutex.Lock()
	balance -= amt
	fmt.Println("withdraw amount ", amt, " and balance", balance)
	mutex.Unlock()
	wg.Done()
}

func main() {
	balance = 1000
	var wg sync.WaitGroup
	wg.Add(2)
	go deposit(1000, &wg)
	go withdraw(1000, &wg)
	wg.Wait()
	fmt.Println("balance :: ", balance)

}

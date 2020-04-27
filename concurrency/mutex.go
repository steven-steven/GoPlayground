package concurrency

import (
	"fmt"
	"sync"
)

var (	//declare 2 variables
	mutex sync.Mutex	//global mutex
	balance int
)

func init(){
	balance = 1000 //initialize
}

func deposit(value int, wg *sync.WaitGroup){
	mutex.Lock()
	fmt.Printf("Depositing %d to account with balance %d\n", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done() //collect itself
}

func withdraw(value int, wg *sync.WaitGroup){
	mutex.Lock()
	fmt.Printf("Withdraing %d to account with balance %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done() //collect itself
}

func TestMutex(){
	var wg sync.WaitGroup
	wg.Add(2)	//wait for two threads
	go withdraw(700, &wg)
	go deposit(500, &wg)
	wg.Wait() //won't proceed till both threads are done

	fmt.Printf("New Balance %d\n", balance)
}
package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func getVal(c chan int) {
	value := rand.Intn(10)	//random int from 1 to 10
	time.Sleep(1000 * time.Millisecond)
	c <-value
}

func TestChannel(){
	valueChannel := make(chan int, 2)	//buffered channel
	defer close(valueChannel)

	go getVal(valueChannel)
	go getVal(valueChannel)

	values := <-valueChannel

	time.Sleep(1000*time.Millisecond)
}
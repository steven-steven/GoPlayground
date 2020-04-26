package tickers

import (
	"fmt"
	"time"
)

func IntervalPrint() {
	//Ticks every 1 second
	ticker:= time.NewTicker(1*time.Second)

	for _ = range ticker.C {
		fmt.Println("tock")
	}
}
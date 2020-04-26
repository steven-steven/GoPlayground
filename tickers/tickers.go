package tickers

import (
	"fmt"
	"time"
)

func IntervalPrint() {
	//Ticker that ticks every 1 second
	ticker:= time.NewTicker(1*time.Second)

	//forEach tick that ticker.C emit value, print.
	for _ = range ticker.C {
		fmt.Println("tock")
	}
}
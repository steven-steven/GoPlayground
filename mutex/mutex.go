package mutex

import (
	"fmt"
	"sync"
)

var (
	mutex sync.mutex
	balance int
)

func init(){
	balance = 1000 //initialize balance to 1000
}
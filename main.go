package main
import (
	"fmt"
	_ "github.com/steven-steven/GoPlayground/consoleInput"
	_ "github.com/steven-steven/GoPlayground/parsing"
	_ "github.com/steven-steven/GoPlayground/sorting"
	_ "github.com/steven-steven/GoPlayground/tickers"
	"github.com/steven-steven/GoPlayground/datastructure"
)

func main(){
	fmt.Println("Hello World")

	//consoleinput.ReadInput()

	//parsing.ReadJson()
	//parsing.ReadXML()

	// sorting.Sort()
	// sorting.CustomSort()

	// go tickers.IntervalPrint()	//run in the background
	// select{}	//so main program runs indefinitely. Otherwise it will stop before first tick

	datastructure.linkedList()
}
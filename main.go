package main
import (
	"fmt"
	_ "github.com/steven-steven/GoPlayground/consoleInput"
	_ "github.com/steven-steven/GoPlayground/parsing"
	_ "github.com/steven-steven/GoPlayground/sorting"
	_ "github.com/steven-steven/GoPlayground/tickers"
	_ "github.com/steven-steven/GoPlayground/datastructure"
	_ "github.com/steven-steven/GoPlayground/concurrency"
	_ "github.com/steven-steven/GoPlayground/filesystem"
	_ "github.com/steven-steven/GoPlayground/graphql"
	"github.com/steven-steven/GoPlayground/restapi"
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

	// datastructure.TraverseLL()

	//concurrency.TestMutex()
	//concurrency.TestChannel()

	// filesystem.GenerateTempFiles()
	// filesystem.UploadFile()

	//graphql.SampleServer1()
	//graphql.SampleServer2()
	//graphql.Test()
}
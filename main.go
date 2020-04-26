package main
import (
	"fmt"
	_ "github.com/steven-steven/GoPlayground/consoleInput"
	_ "github.com/steven-steven/GoPlayground/parsing"
	"github.com/steven-steven/GoPlayground/sorting"
)

func main(){
	fmt.Println("Hello World")
	//consoleinput.ReadInput()
	//parsing.ReadJson()
	//parsing.ReadXML()
	sorting.Sort()
	sorting.CustomSort()
}
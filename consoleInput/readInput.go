package consoleinput

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInput(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("SimpleShell")
	fmt.Println("-----------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')	//read string until \n
		text = strings.Replace(text, "\n", "", -1)	//strings library to replace substring
		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello yourself")
		}
	}
}
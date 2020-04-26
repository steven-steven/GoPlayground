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

	//or use scanner
	for {
		fmt.Print("-> ")
		readChar(reader)
	}
}

func readLine(reader *bufio.Reader){
	text, _ := reader.ReadString('\n')	//read string until \n
	text = strings.Replace(text, "\r\n", "", -1)	//strings library to replace substring
	if strings.Compare("hi", text) == 0 {
		fmt.Println("hello yourself")
	}
}

func readChar(reader *bufio.Reader){
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	
	// print out the unicode value i.e. A -> 65, a -> 97
	fmt.Println(char)
	
	switch char {
		case 'A':
			fmt.Println("A Key Pressed")
			break
		case 'a':
			fmt.Println("a Key Pressed")
			break
	}
}
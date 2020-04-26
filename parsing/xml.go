package parsing

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/xml"
)

type Users_2 struct {
	XMLName xml.Name `xml:"users"`	//name of this tag
	Users   []User_2   `xml:"user"`		//tags that it contain
}

type User_2 struct {
	XMLName xml.Name `xml:"user"`	//name of this tag
	Type    string   `xml:"type,attr"`	//attribute of <User>, with name type
	Name    string   `xml:"name"`	//contains a tag <name>
	Social  Social_2   `xml:"social"`	//contains a tag <social>
}

type Social_2 struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}

func ReadXML(){
	// Open our xmlFile
	xmlFile, err := os.Open("parsing/users.xml")
	if err != nil {
			fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.xml")
	defer xmlFile.Close()

	// io.reader to []byte
	byteValue, _ := ioutil.ReadAll(xmlFile)

	var users Users_2
	xml.Unmarshal(byteValue, &users)

	//iterate through users
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}
}
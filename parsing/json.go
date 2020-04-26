package parsing

import (
	"fmt"
	"os"	//to open file
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func ReadJson(){
	//open file -> convert to a byte array -> unmarshal json to struct
	// Open our jsonFile (io.Reader)
	jsonFile, err := os.Open("users.json")
	if err != nil {
			fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	//------Method 1---------------------------------------
	//Parsing with Structs -if you know the structure
	byteValue, _ := ioutil.ReadAll(jsonFile)	//io.reader to []byte
	var users Users
	json.Unmarshal(byteValue, &users)

	//print results
	for i := 0; i < len(users.Users); i++ {
    fmt.Println("User Type: " + users.Users[i].Type)
    fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))	//int to string (opposite is atoi string to int)
    fmt.Println("User Name: " + users.Users[i].Name)
    fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}

	//------Method 2---------------------------------------
	//Parse with map of string to anything (when working with unstructured data)
	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)

	fmt.Println(result["users"])
}
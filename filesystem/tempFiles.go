package filesystem

//create globally unique files/directories
//we don't care about names
import (
	"fmt"
	"io/ioutil"
	"os"
)

func GenerateTempFiles() {
	//generate temp directory
	tempDir, err := ioutil.TempDir("filesystem/", "cars-")
	if err != nil {
			fmt.Println(err)
	}
	defer os.RemoveAll(tempDir)

	// ioutil.TempFile returns either a file or error.
	// car-images is the file's . We will create picture
	file, err := ioutil.TempFile(tempDir, "car-*.png")
	if err != nil {
			fmt.Println(err)
	}
	defer os.Remove(file.Name())	//delete the file at end
	defer file.Close()

	//file Write to png file with string/[]byte (is this idiomatic way to use if for error checking)
	if _, err := file.Write([]byte("hello world\n")); err != nil {
		fmt.Println(err)
	}

	//read and print contents
	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
			fmt.Println(err)
	}
	fmt.Print(string(data))
}
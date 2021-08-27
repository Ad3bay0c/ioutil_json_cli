package files

import (
	// "fmt"
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"

	// "io"
	"os"
	// "time"
)

func File() {
	// out, _ := os.Create("files.txt")
	// defer out.Close()

	// resp, _ := http.Get("http://www.gutenberg.org/files/18581/18581.txt")
	// defer resp.Body.Close()

	// n, _ := io.Copy(out, resp.Body)
	// fmt.Println(n)

	// Create Directory, Remove, and Rename

	// os.Mkdir("newFile", 0777)

	// fmt.Println("Created Successfully")

	// // os.Remove("newFile")
	// os.RemoveAll("files/newFile")
	// os.Rename("newFile", "oldFile")
	// fmt.Println("removed Successfully")

	// os.MkdirAll("files/newFile", 0777)
	// time.Sleep(time.Second * 1)

	//Create a File and Write to it

	// file, err := os.Create("file.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// for i:=0; i<10; i++ {
	// 	file.WriteString("We move and let's go\r\n")
	// 	file.Write([]byte("We move and let's go\r\n"))
	// }

	fmt.Println(os.Getwd())

	file, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Seek(6, 1)
	buf := make([]byte, 1024)
	file.Read(buf)

	fmt.Println(string(buf))

	// file.WriteString("Okay, Another one\r\n")

	// Using os package
	// buf := make([]byte, 1024)
	// for {
	// 	n, _ := file.Read(buf)
	// 	if n == 0 {
	// 		break
	// 	}
	// 	os.Stdout.Write(buf)
	// }
	// Using bufio package
	// inputReader := bufio.NewReader(file)
	// for {
	// 	input, err := inputReader.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	os.Stdout.WriteString(input)
	// }

	file, _ = os.OpenFile("file.txt", os.O_WRONLY, 0666)
	outputWriter := bufio.NewWriter(file)

	// outputWriter.WriteString("Again we made it\r\n")
	outputWriter.WriteString("Let's try something new\r\n")
	outputWriter.WriteString("No matter how strong you are\r\n")
	outputWriter.WriteString("We gtaher dey now\r\n")
	outputWriter.WriteString("Okay Thanks to your\n")
	outputWriter.WriteString("{\"users\":[{\"id\":1629439632,\"name\":\"Ade\",\"occupation\":\"Programmer\",\"phone\":1234567},{\"id\":1629439632,\"name\":\"Adebayo\",\"occupation\":\"Gopher\",\"phone\":36192765},{\"id\":1629439632,\"name\":\"Clinton\",\"occupation\":\"Gopher Json\",\"phone\":36192765}]}	\n")
	outputWriter.Flush()
	// outputWriter.WriteString("Okay, WOnderfully made")

	r := strings.NewReader("We took it where we needed to take it to")
	res, _ := ioutil.ReadAll(r)
	fmt.Printf("%s\n", res)

	newRead, err := ioutil.ReadFile("file.csv")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", newRead)

	// data := []byte("{\"users\":[{\"id\":1629439632,\"name\":\"Ade\",\"occupation\":\"Programmer\",\"phone\":1234567},{\"id\":1629439632,\"name\":\"Adebayo\",\"occupation\":\"Gopher\",\"phone\":36192765}")
	// error := ioutil.WriteFile("file.csv", data, 0666)
	// if error != nil {
	// 	panic(error)
	// }
	// data = []byte("Okay just testing if it is gonna be cleaned up and updated")
	// ioutil.WriteFile("file.txt", data, 0666)

}

package io_package

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func IoTesting() {
	r := strings.NewReader("Okay testing IO Package\n")

	_, e := io.Copy(os.Stdout, r)

	if e != nil {
		fmt.Println(e)
	}

	// CopyBuffer
	r = strings.NewReader("This is testing for Buffer")
	var buf = make([]byte, 10)
	_, er := io.CopyBuffer(os.Stdout, r, buf)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(string(buf))

	r= strings.NewReader("We Move to Read AtLeast")
	// var buff = make([]byte, 10)

	b, err := io.ReadAll(r)
	// _,err := io.ReadAtLeast(r, buff, 10)
	if err != nil {
		fmt.Println(er)
	}
	fmt.Println(string(b))


}
func IoPipeTesting() {
	r, w := io.Pipe()
	go func ()  {
		fmt.Fprint(w, "Okay Good\t")
		fmt.Fprint(w, "Okay Good\t")
		fmt.Fprint(w, "Okay Good\t")
		fmt.Fprint(w, "Okay Good\n")
		w.Close()
	}()

	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		fmt.Println(err)
	}

	io.WriteString(os.Stdout, "Writing directly to string using io.Writestring\n")
}

func InputReader() {
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Print("Please Enter your name: ")
	input, _ := inputReader.ReadString('\n')
	fmt.Println("Your Name is: ", input)
}

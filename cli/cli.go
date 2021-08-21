package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

const FILENAME = "file.csv"

type List struct {
	Content string
	Done    bool
	Date    int64
}

type Lists struct {
	Lists []List
}

var Lst = Lists{}

func init() {
	loadTodoList()
}

func loadTodoList() {
	file, err := ioutil.ReadFile("file.csv")
	if err != nil {
		panic(err)
	}
	
	_ = json.Unmarshal(file, &Lst)
}

func (lists Lists) All() {
	var i = 0
	if len(lists.Lists) == 0 {
		fmt.Println("Empty List, Please Add Content")
	}

	for _, val := range lists.Lists {
		if val.Done {
			continue
		}
		i += 1
		fmt.Printf("%d. %s - %v\n", i, val.Content, time.Unix(val.Date, 64).Format("Mon, 02 Jan 2006 15:04:05"))
	}
}

func (lists *Lists) AddContent(content string) {
	list := List{Content: content, Date: time.Now().Unix()}
	lists.Lists = append(lists.Lists, list)

	addToFile()
}
func addToFile() {
	res, _ := json.Marshal(Lst)

	_ = ioutil.WriteFile(FILENAME, res, 0666)
}

func (lists *Lists) Done(value int) {
	if len(lists.Lists) < value {
		fmt.Println("Please enter a valid no")
	} else {
		Lst.Lists[value-1].Done = true
		addToFile()
	}
}

func (lists *Lists) Undone(value int) {
	if len(lists.Lists) < value {
		fmt.Println("Please enter a valid no")
	} else {
		Lst.Lists[value-1].Done = false
		addToFile()
	}
}

func (lists *Lists) Cleanup() {
	for idx, val := range lists.Lists {
		if val.Done && idx < len(lists.Lists)-1 {
			lists.Lists = append(lists.Lists[:idx], lists.Lists[idx+1:]...)
		} else if val.Done && idx == len(lists.Lists)-1{
			lists.Lists = lists.Lists[:idx]
		}
	}
	addToFile()
}

// func CliTest() {

// 	// inputReader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Please enter any of the following to perform a function or 0 to exit")
// 	fmt.Printf("1.\t Lists all Todo\n2.\t Add To Todo\n3.\t Done Todo\n4.\t Undone Todo\n5.\t Cleanup All Done\n0.\t Exit\n-1.\t Show Menu\n Please Enter a Number:  ")

// 	// input, err:= inputReader.ReadString('\n')
// 	var inp int
// 	fmt.Scanf("%d", &inp)
// 	for {
// 		switch inp {
// 			case 0:
// 				return
// 			case 1:
// 				Lst.All()
// 				continue
// 			case 2:
// 				var content string
// 				fmt.Print("Please Enter the content you wanna Add: ")
// 				fmt.Scanf("%s", &content)
// 				Lst.AddContent(content)
// 				continue
// 			case 3:
// 				var content int
// 				fmt.Print("Please Enter the No you wanna Done: ")
// 				fmt.Scanf("%d", &content)
// 				Lst.Done(content)
// 				continue
// 			case 4:
// 				var content int
// 				fmt.Print("Please Enter the No you wanna Undone: ")
// 				fmt.Scanf("%s", &content)
// 				Lst.Undone(content)
// 				continue
// 			case 5:
// 				Lst.Cleanup()
// 				continue
// 			default:
// 				fmt.Printf("1.\t Lists all Todo\n2.\t Add To Todo\n3.\t Done Todo\n4.\t Undone Todo\n5.\t Cleanup All Done\n0.\t Exit\n-1.\t Show Menu\n Please Enter a Number:  ")
// 				continue

// 			}
// 	}
	

	

// 	// fmt.Printf("%T",inp)
// }

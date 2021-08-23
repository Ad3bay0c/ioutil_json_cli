package model

import (
	"encoding/json"
	"github.com/Ad3bay0c/week4Curriculum/cli"
	"io/ioutil"
)

var Lst = cli.Lists{}

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

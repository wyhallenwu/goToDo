package file

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"todotxt/internal/entry"
)

func init() {
	CreateDir()
	CreateFile()
}

// CreateDir is to create a directory for todoList
func CreateDir() {
	dir := "../todoList"
	err := os.Mkdir(dir, os.FileMode(0777))
	if os.IsExist(err) {
		log.Println("directory ", dir, " exists")
	}
}

// CreateFile is to create two json files
func CreateFile() {
	fileTodo, err := os.Create("../todoList/ToDo.json")
	defer fileTodo.Close()
	if os.IsExist(err) {
		log.Println("ToDo.json already exists")
	}
	fileDone, err := os.Create("../todoList/Done.json")
	defer fileDone.Close()
	if os.IsExist(err) {
		log.Println("Done.json already exists")
	}
}

func ReadTodoList() []entry.Entry {
	path := "../todoList/ToDo,json"
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Println("ReadTodoList break up")
	}
	var entryList []entry.Entry
	decoder := json.NewDecoder(file)
	err = decoder.Decode(entryList)
	if err != nil {
		log.Println("docoder wrong: ", err.Error())
	}
	return entryList
}

func Insert2Todo(entryItem *entry.Entry) {
	_, err := json.Marshal(*entryItem)
	if err != nil {
		log.Println(err.Error())
	}
	//	TODO
}

func Run() {
	fmt.Println("initialize successfully")
}

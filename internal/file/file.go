package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"todotxt/internal/entry"
)

// Initialize must run when first using this app
func Initialize() {
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

// Insert2File is to insert an entry to the aiming file
func Insert2File(entryItem *entry.Entry, filename string) {
	entryList := ReadFile(filename)
	entryList = append(entryList, *entryItem)
	data, _ := json.MarshalIndent(entryList, "", "    ")
	err := ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		log.Println(err)
	}
}

func Write2File(entryList []entry.Entry, filename string) {
	data, _ := json.MarshalIndent(entryList, "", "    ")
	err := ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		log.Println(err)
	}
}

func CreateAndInsert(description string) {
	en := entry.NewEntry(description)
	InsertEntry(en)
}

// InsertEntry is to insert an entry into the aiming file
func InsertEntry(inEn *entry.Entry) {
	// entry is newly created
	if !inEn.GetterStatus() {
		Insert2File(inEn, "../todoList/ToDo.json")
	} else {
		Insert2File(inEn, "../todoList/Done.json")
	}
}

func ReadFile(filename string) []entry.Entry {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// read json file
	entryList := make([]entry.Entry, 0)
	err = json.Unmarshal(content, &entryList)
	return entryList
}

func Run() {
	fmt.Println("initialize successfully")
}

func ShowFile(filename string) {
	entryList := ReadFile(filename)
	for ix, en := range entryList {
		en.ShowEntry(ix + 1)
	}
}

func RemoveEntry(index int, filename string) *entry.Entry {
	entryList := ReadFile("../todoList/ToDo.json")
	for ix, _ := range entryList {
		if (ix + 1) == index {
			reEn := entryList[ix]
			entryList = append(entryList[:ix], entryList[ix+1:]...)
			Write2File(entryList, filename)
			return &reEn
		}
	}
	return nil
}

func Finished(index int) {
	delEn := RemoveEntry(index, "../todoList/ToDo.json")
	delEn.SetStatus(true)
	Insert2File(delEn, "../todoList/Done.json")
}

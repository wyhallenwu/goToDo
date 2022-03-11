package file

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"todotxt/internal/entry"
)

//initialize ======================================

// Initialize must run when first using this app
func Initialize() {
	InitDir()
	InitFile()
	log.Println("initialize successfully")
}

// InitDir is to create a directory for todoList
func InitDir() {
	dir := "../todoList"
	err := os.Mkdir(dir, os.FileMode(0777))
	if os.IsExist(err) {
		log.Println("directory ", dir, " exists")
	}
}

// InitFile is to create two json files
func InitFile() {
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

// file operations=====================================

// ReadFile is to read a json file and return an []Entry
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

func CreateAndInsertEntry(description string) {
	en := entry.NewEntry(description)
	InsertEntryToFile(en)
}

// InsertEntryToFile is to insert an entry into the aiming file
func InsertEntryToFile(inEn *entry.Entry) {
	// entry is newly created
	if !inEn.GetterStatus() {
		InsertFile(inEn, "../todoList/ToDo.json")
	} else {
		InsertFile(inEn, "../todoList/Done.json")
	}
}

// InsertFile is to insert an entry to the aiming file
func InsertFile(entryItem *entry.Entry, filename string) {
	entryList := ReadFile(filename)
	entryList = append(entryList, *entryItem)
	data, _ := json.MarshalIndent(entryList, "", "    ")
	err := ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		log.Println(err)
	}
}

func WriteFile(entryList []entry.Entry, filename string) {
	data, _ := json.MarshalIndent(entryList, "", "    ")
	err := ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		log.Println(err)
	}
}

// ShowFile is to format all entries and show in the terminal
func ShowFile(filename string) {
	entryList := ReadFile(filename)
	for ix, en := range entryList {
		en.ShowEntry(ix + 1)
	}
}

// EntryDone changes entry's status to done and do subsequent operations
func EntryDone(index int) {
	delEn := RemoveEntry(index, "../todoList/ToDo.json")
	delEn.SetStatus(true)
	InsertFile(delEn, "../todoList/Done.json")
}

// RemoveEntry removes an entry by searching index
func RemoveEntry(index int, filename string) *entry.Entry {
	entryList := ReadFile("../todoList/ToDo.json")
	for ix, _ := range entryList {
		if (ix + 1) == index {
			reEn := entryList[ix]
			entryList = append(entryList[:ix], entryList[ix+1:]...)
			WriteFile(entryList, filename)
			return &reEn
		}
	}
	return nil
}

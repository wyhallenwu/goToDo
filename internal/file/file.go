package file

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"todotxt/internal/entry"
)

//initialize ======================================
const (
	dir      = "/home/wuyuheng/Desktop/todoList"
	TodoFile = dir + "/ToDo.json"
	DoneFile = dir + "/Done.json"
)

// Initialize must run when first using this app
func Initialize() {
	InitDir()
	InitFile()
	log.Println("initialize successfully")
}

// InitDir is to create a directory for todoList
func InitDir() {
	err := os.Mkdir(dir, os.FileMode(0777))
	if os.IsExist(err) {
		log.Println("directory ", dir, " exists")
	}
}

// InitFile is to create two json files
func InitFile() {
	fileTodo, err := os.Create(TodoFile)
	defer fileTodo.Close()
	if os.IsExist(err) {
		log.Println("ToDo.json already exists")
	}
	fileDone, err := os.Create(DoneFile)
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
		InsertFile(inEn, TodoFile)
	} else {
		InsertFile(inEn, DoneFile)
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
		en.PrintEntry(ix + 1)
	}
}

// EntryDone changes entry's status to done and do subsequent operations
func EntryDone(index int) {
	delEn := RemoveEntry(index, TodoFile)
	delEn.SetStatus(true)
	InsertFile(delEn, DoneFile)
	log.Println(index, " done")
}

// RemoveEntry removes an entry by searching index
func RemoveEntry(index int, filename string) *entry.Entry {
	entryList := ReadFile(TodoFile)
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

func GroupByProject(filename string, project string) []entry.Entry {
	entryList := ReadFile(filename)
	groupEntry := make([]entry.Entry, 0)
	for _, item := range entryList {
		if item.Project == project {
			groupEntry = append(groupEntry, item)
		}
	}
	return groupEntry
}

// PrintGroup prints all items which belongs to project
func PrintGroup(filename string, project string) {
	groupEntry := GroupByProject(filename, project)
	for ix, item := range groupEntry {
		item.PrintEntry(ix)
	}
}

// AddProjectToItem changes the group the item belongs to
func AddProjectToItem(index int, project string, filename string) {
	entryList := ReadFile(filename)
	changeEntry := &entryList[index]
	changeEntry.SetProject(project)
	// write back to file
	WriteFile(entryList, TodoFile)
}

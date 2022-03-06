package file

import (
	"log"
	"os"
)

func init() {
	CreateDir()
	CreateFile()
}

func CreateDir() {
	dir := "./todoList"
	err := os.Mkdir(dir, os.FileMode(0777))
	if os.IsExist(err) {
		log.Println("directory ", dir, " exists")
	} else {
		log.Println(err)
	}
}

func CreateFile() {
	fileTodo, err := os.Create("./todoList/ToDo.txt")
	defer fileTodo.Close()
	if os.IsExist(err) {
		log.Println("ToDo.txt already exists")
	} else {
		log.Println(err)
	}
	fileDone, err := os.Create("./todoList/Done.txt")
	defer fileDone.Close()
	if os.IsExist(err) {
		log.Println("Done.txt already exists")
	} else {
		log.Println(err)
	}
}

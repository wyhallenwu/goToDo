package main

import (
	"fmt"
	"todotxt/internal/file"
)

func main() {
	file.Run()
	file.CreateAndInsert("test insert")
	file.Finished(1)
	file.ShowFile("../todoList/Done.json")
	fmt.Println("=======================")
	file.ShowFile("../todoList/ToDo.json")
}

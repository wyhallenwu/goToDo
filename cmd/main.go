package main

import (
	"flag"
	"os"
	"todotxt/internal/file"
)

var (
	initial   bool   // initial is to initialize all files when first use
	add       string // add represents description
	list      bool   // list to_do file
	show      bool   // show done file
	doneIndex int    // doneIndex represents entry is done
	help      bool   // show help info
)

func init() {
	flag.BoolVar(&initial, "init", false, "initialize all files")
	flag.StringVar(&add, "a", "", "add descriptions of the entry")
	flag.BoolVar(&list, "l", false, "list all todo entries")
	flag.BoolVar(&show, "s", false, "show add done entries")
	flag.IntVar(&doneIndex, "d", 0, "entry doneIndex is finished")
	flag.BoolVar(&help, "h", false, "show help information")
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}
	if initial {
		file.Initialize()
	}
	if add != "" {
		file.CreateAndInsertEntry(add)
	}
	if list {
		file.ShowFile("../todoList/ToDo.json")
	}
	if show {
		file.ShowFile("../todoList/Done.json")
	}
	if doneIndex > 0 {
		file.EntryDone(doneIndex)
	}
}

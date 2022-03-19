package main

import (
	"flag"
	"os"
	"todotxt/internal/file"
)

var (
	initial     bool   // initial is to initialize all files when first use
	add         string // add represents description
	list        bool   // list to_do file
	show        bool   // show done file
	index       int    // index represent the item index
	done        bool
	group       string // category
	help        bool   // print help info
	changeGroup bool   // changeGroup
)

func init() {
	flag.BoolVar(&initial, "init", false, "initialize all files")
	flag.StringVar(&add, "a", "", "add descriptions of the entry")
	flag.BoolVar(&list, "l", false, "list all todo entries")
	flag.BoolVar(&show, "s", false, "show add done entries")
	flag.IntVar(&index, "i", 0, "entry index")
	flag.BoolVar(&done, "d", false, "done or not")
	flag.BoolVar(&help, "h", false, "show help information")
	flag.StringVar(&group, "p", "", "please input project category")
	// todo: add change group
	flag.BoolVar(&changeGroup, "c", false, "change or not")
}

func main() {
	flag.Parse()
	// parse command and run
	switch {
	case help:
		flag.Usage()
		os.Exit(0)
	case initial:
		file.Initialize()
	case add != "":
		file.CreateAndInsertEntry(add)
	case list:
		file.ShowFile(file.TodoFile)
	case show:
		file.ShowFile(file.DoneFile)
	case done && index > 0:
		file.EntryDone(index)
	case changeGroup && index > 0 && group != "": // todo: bugs exists
		file.AddProjectToItem(index, group, file.TodoFile)
	case group != "":
		file.PrintGroup(file.TodoFile, group)
	}
}

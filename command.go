package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := &CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specified by the title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index and specify a new title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo by its index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo by its index")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Del != -1:
		todos.delete(cf.Del)
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid format for edit. Use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid index for edit.")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	default:
		fmt.Println("No valid flags provided. Use -h for help.")
		os.Exit(1)
	}
}

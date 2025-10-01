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

func newCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "add a new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "edit todo by intex and specify a new title")
	flag.IntVar(&cf.Del, "del", -1, "specify todo by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "specify todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all todo")

	flag.Parse()

	return &cf

}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Erro, invalid format for edit, Pleae us id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Errors: invalid index for edit")
			os.Exit(1)
		}
		todos.edit(index, parts[1])
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != -1:
		todos.delete(cf.Del)
	default:
		fmt.Println("Invalid Commad")
	}

}

package main

// import "fmt"

func main() {
	todos := Todos{}
	storage := NewStrorage[Todos]("todos.json")

	storage.Load(&todos)
	cmdFlags := newCmdFlags()

	cmdFlags.Execute(&todos)
	// todos.toggle(0)
	storage.Save(todos)
	// todos.toggle(0)

	todos.print()
}

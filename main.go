package main

func main() {

	todos := Todos{}
	Storage := NewStorage[Todos]("todos.json")

	Storage.Load(&todos)

	todos.print()

	Storage.Save(todos)
}

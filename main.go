package main

func main() {

	todos := Todos{}

	todos.add("Buy Milk")
	todos.add("Buy bread")

	todos.toggle(1)

	todos.print()
}

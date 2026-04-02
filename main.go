package main

import "fmt"

func main() {

	todos := Todos{}

	todos.add("Buy Milk")
	todos.add("Buy bread")

	fmt.Printf("%+v\n\n", todos) // %+v is used to print the struct with field names and values.

	todos.delete(0)

	fmt.Printf("%+v\n\n", todos)
}

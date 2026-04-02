package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

// declaring a struct for our todo item.
type Todo struct {
	Title        string
	Completed    bool
	Created_at   time.Time
	Completed_at *time.Time // Completed at is a pointer reference because when creating the struct instance, the completed at can be null.
}

// Contains all our Todos
// We are creating a slice of Todo heare, because we want to be able to add new Todo items and slices allow us to create methods attached to it.
type Todos []Todo

// Creating Add method
func (todos *Todos) add(title string) { // (todos *Todos) is the receiver of the method so that we can call this method on the Todos struct. The * is used to indicate that we are passing a pointer reference to the Todos struct, so that we can modify the original struct instance when we call this method.
	todo := Todo{
		Title:        title,
		Completed:    false,
		Completed_at: nil,
		Created_at:   time.Now(),
	}

	*todos = append(*todos, todo) // The new todo is added to the original Todos struct instance by dereferencing the pointer reference and appending the new todo to the slice. The original Todos struct instance being Todos{} in main.go.
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos // Creates a local copy t of the dereferenced slice. This is done to work with the slice without repeatedly dereferencing the pointer.

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...) // This line modifies the original Todos struct instance by slicing the local copy t to exclude the element at the specified index and then appending the remaining elements back to the original Todos slice. The ... is used to unpack the elements of the sliced portion of t so that they can be appended correctly to the original Todos slice.

	return nil
}

func (todos *Todos) toggle(index int) error {

	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		now := time.Now()
		t[index].Completed_at = &now
	} else {
		t[index].Completed_at = nil
	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit(index int, title string) error {

	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil
}

func (todos *Todos) print() {

	table := table.New(os.Stdout)

	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, todo := range *todos {
		completed := "❌"
		completedAt := ""

		if todo.Completed {
			completed = "✅"
			completedAt = todo.Completed_at.Format(time.RFC1123)
		}

		table.AddRow(strconv.Itoa(index), todo.Title, completed, todo.Created_at.Format(time.RFC1123), completedAt)
	}

	table.Render()

}

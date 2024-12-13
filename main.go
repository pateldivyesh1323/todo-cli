package main

import (
	"fmt"
	"os"
	"todo-cli/cmd"
	"todo-cli/storage"
	"todo-cli/todo"
)

const storageFile = "todos.json"

var todoManager *todo.TodoManager

func init() {
	// Create new file if it doesn't exist
	_, err := os.Open(storageFile)
	if os.IsNotExist(err) {
		_, err := os.Create(storageFile)
		fmt.Println("✅ Storage file created for storing todos")
		if err != nil {
			fmt.Println("Error creating storage file:", err)
		}
	} else if err != nil {
		fmt.Println("Error opening storage file:", err)
	}

	// Load todos from storage
	todos, err := storage.LoadTodos(storageFile)
	if err != nil {
		fmt.Println("No todos found. Start adding todos with `todo add <todo>`")
	}
	todoManager = todo.NewTodoManager(todos)
}

func main() {
	defer func() {
		if err := storage.SaveTodos(storageFile, todoManager.GetTodos()); err != nil {
			fmt.Println("Error saving todos:", err)
		}
	}()
	cmd.Execute()
}

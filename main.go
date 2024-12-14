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
			panic("Error creating storage file: " + err.Error())
		}
	} else if err != nil {
		panic("Error opening storage file: " + err.Error())
	}

	// Load todos from storage
	todos, _ := storage.LoadTodos(storageFile)
	todoManager = todo.NewTodoManager(todos)
	cmd.SetTodoManager(todoManager)
}

func main() {
	defer func() {
		if err := storage.SaveTodos(storageFile, todoManager.GetTodos()); err != nil {
			fmt.Println("Error saving todos:", err)
		}
	}()
	cmd.Execute()
}

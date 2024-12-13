package storage

import (
	"encoding/json"
	"os"
	"todo-cli/todo"
)

func SaveTodos(fileName string, todos []todo.Todo) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(todos)
}

func LoadTodos(fileName string) ([]todo.Todo, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var todos []todo.Todo
	err = json.NewDecoder(file).Decode(&todos)
	return todos, err
}

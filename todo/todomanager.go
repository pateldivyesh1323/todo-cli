package todo

import (
	"errors"
	"fmt"
	"time"
)

type TodoManager struct {
	todos  []Todo
	nextId int
}

func NewTodoManager(loadedTodo []Todo) *TodoManager {
	loadedNextId := 1
	if len(loadedTodo) > 0 {
		loadedNextId = loadedTodo[len(loadedTodo)-1].ID + 1
	}
	return &TodoManager{todos: loadedTodo, nextId: loadedNextId}
}

func (tm *TodoManager) GetTodos() []Todo {
	return tm.todos
}

func (tm *TodoManager) AddTodo(title string) {
	todo := Todo{ID: tm.nextId, Title: title, Completed: false, CreatedAt: time.Now()}
	tm.todos = append(tm.todos, todo)
	tm.nextId++
}

func (tm *TodoManager) ListTodos() {
	if len(tm.todos) == 0 {
		fmt.Println("No todos found. Start adding todos with `todo add <todo>`")
		return
	}
	for _, todo := range tm.todos {
		status := "❌"
		if todo.Completed {
			status = "✅"
		}
		fmt.Printf("%d: %s [%s]\n", todo.ID, todo.Title, status)
	}
}

func (tm *TodoManager) MarkCompleted(id int) error {
	for _, todo := range tm.todos {
		if todo.ID == id {
			tm.todos[id].Completed = true
			return nil
		}
	}
	return errors.New("Todo not found")
}

func (tm *TodoManager) Delete(id int) error {
	for i, todo := range tm.todos {
		if todo.ID == id {
			tm.todos = append(tm.todos[:i], tm.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("Todo not found")
}

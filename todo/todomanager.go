package todo

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
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

	println("-----------------------------")
	println("**** List of Your Todos ****")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Status"})
	for _, todo := range tm.todos {
		status := "Incomplete"
		if todo.Completed {
			status = "Completed"
		}
		table.Append([]string{
			strconv.Itoa(todo.ID),
			todo.Title,
			status,
		})
	}
	table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	table.SetCenterSeparator("|")
	table.SetRowSeparator("-")
	table.SetColumnSeparator("|")
	table.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	table.Render()
}

func (tm *TodoManager) MarkCompleted(id int) error {
	for i, t := range tm.todos {
		if t.ID == id {
			t.Completed = true
			tm.todos = append(tm.todos[:i], append([]Todo{t}, tm.todos[i+1:]...)...)
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

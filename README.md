# Todo CLI Application

This is a simple command-line To-Do application written in Go. The application allows users to manage their tasks efficiently by providing commands to add, list, delete and mark tasks as completed. Tasks are displayed in a neatly formatted table with borders.

---

## Features

-   Add new tasks with a title.
-   List all tasks in a tabular format.
-   Mark tasks as completed.
-   Delete tasks
-   Tasks are stored in file.

---

## Prerequisites

-   [Go](https://go.dev/) (version 1.18 or higher installed).
-   A terminal or command prompt to run the application.

---

## Installation

1. Clone the repository:

    ```bash
    git clone <repository-url>
    cd todo-cli
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Build the project:

    ```bash
    go build -o ./bin/todo.exe
    ```

4. Run the application:
    ```bash
    ./bin/todo
    ```

---

## Usage

### Add a Task

To add a new task, use the following command:

```bash
todo add "Task Title"
```

Example:

```bash
todo add "Buy groceries"
```

### List All Tasks

To list all tasks in a tabular format, use:

```bash
todo list
```

Output example:

```
----------------------------
**** List of Your Todos ****
+----+----------------------+-----------+
| ID |        TITLE         |  STATUS   |
+----+----------------------+-----------+
| 1  | Buy groceries        | Incomplete|
| 2  | Write documentation  | Completed |
| 3  | Clean the house      | Incomplete|
+----+----------------------+-----------+
```

### Mark a Task as Completed

To mark a task as completed, use:

```bash
todo done <task-id>
```

Example:

```bash
todo done 1
```

### Delete a task

To delete a task, use:

```bash
todo delete <task-id>
```

Example:

```bash
todo delete 1
```

### Help

For help with available commands, run:

```bash
todo help
```

---

## Example Execution

```bash
$ todo add "Buy groceries"
Task added successfully!

$ todo list
----------------------------
**** List of Your Todos ****
+----+----------------------+-----------+
| ID |        TITLE         |  STATUS   |
+----+----------------------+-----------+
| 1  | Buy groceries        | Incomplete|
+----+----------------------+-----------+

$ todo done 1
Task marked as completed!

$ todo list
----------------------------
**** List of Your Todos ****
+----+----------------------+-----------+
| ID |        TITLE         |  STATUS   |
+----+----------------------+-----------+
| 1  | Buy groceries        | Completed |
+----+----------------------+-----------+
```

---

## Future Enhancements

-   Persist tasks to a file or database.
-   Add task priorities (e.g., High, Medium, Low).
-   Implement task deletion.
-   Add filtering options (e.g., show only completed tasks).
-   Allow editing of tasks.

---

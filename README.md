# Go ToDo CLI

<p>A simple command-line ToDo list application written in Go. This tool allows you to manage your daily tasks using a CLI interface powered by Cobra.</p>

Made for Task Tracker CLI on [roadmap.sh](https://roadmap.sh/projects/task-tracker)

## Features

- **Interactive CLI**: Powered by Cobra for a smooth command-line experience.
- **Task Management**:
  - **Add**: Create new tasks easily.
  - **List**: View all your tasks, including their status.
  - **Update**: Modify existing tasks.
  - **Complete**: Mark tasks as done.
  - **Remove**: Delete tasks you no longer need.
  - **Clear**: Remove all tasks at once.
- **Persistence**: Automatically saves your tasks to a JSON file.

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) installed on your machine.

### Installation

1. Clone the repository:
   ```bash
   git clone <https://github.com/condog220/ToDo.git>
   ```

2. Navigate to the directory:
   ```bash
   cd ToDo
   ```

## Usage

You can run the application directly using `go run main.go` followed by the command.

### Add a Task
Add a new task to your list.
```bash
go run main.go add --task "Buy groceries"
# or use the shorthand flag
go run main.go add -t "Buy groceries"
```

### List Tasks
View all tasks in your list.
```bash
go run main.go list
```

### Update a Task
Update the description of an existing task.
```bash
# Update task at index 1
go run main.go update --index 1 --task "Buy organic groceries"
# or
go run main.go update -i 1 -t "Buy organic groceries"
```

### Mark Task as Completed
Mark a task as done.
```bash
# Complete task at index 1
go run main.go markCompleted --index 1
# or
go run main.go markCompleted -i 1
```

### Remove a Task
Delete a task from the list.
```bash
# Remove task at index 1
go run main.go remove --index 1
# or
go run main.go remove -i 1
```

### Clear List
Delete all tasks from the list.
```bash
go run main.go clearList
```

## Project Structure

- `main.go`: Entry point, initializes the Cobra CLI.
- `cmd/`: Contains all CLI command definitions (`add`, `list`, `update`, etc.).
- `cmd/todo.go`: Core logic and data structures.
- `cmd/json.go`: Handles JSON file operations.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var toDos tasks
	var choice int

	loadedTasks, err := LoadfromJson()
	if err != nil {
		fmt.Println("Error loading tasks from JSON:", err)
	}
	toDos = loadedTasks

	for choice != 7 {
		printInstructions()
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var task string
			fmt.Println("Enter the task to add:")
			reader := bufio.NewReader(os.Stdin)
			task, _ = reader.ReadString('\n')
			toDos.add(task)
		case 2:
			var index int
			fmt.Println("Enter the task number to remove:")
			fmt.Scanln(&index)
			toDos.remove(index - 1)
		case 3:
			var index int
			var task string
			fmt.Println("Enter the task number to update:")
			fmt.Scanln(&index)
			fmt.Println("Enter the new task description:")
			reader := bufio.NewReader(os.Stdin)
			task, _ = reader.ReadString('\n')
			toDos.update(index-1, task)
		case 4:
			var index int
			fmt.Println("Enter the task number to mark as completed:")
			fmt.Scanln(&index)
			toDos.markCompleted(index - 1)
		case 5:
			toDos.list()
		case 6:
			err := SavetoJson(toDos)
			fmt.Println("Tasks saved to JSON successfully!")
			if err != nil {
				fmt.Println("Error saving to JSON:", err)
			}
		case 7:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}

func printInstructions() {
	fmt.Println("Welcome to a To Do List application")
	fmt.Println("--------------------------------")
	fmt.Println("1. Add Task")
	fmt.Println("2. Remove Task")
	fmt.Println("3. Update Task")
	fmt.Println("4. Mark Task as Completed")
	fmt.Println("5. List Tasks")
	fmt.Println("6. Save tasks to JSON")
	fmt.Println("7. Exit")
	fmt.Println("Enter your choice:")
}

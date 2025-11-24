package cmd

import (
	"fmt"
	"strings"
	"time"
)

type Task struct {
	Task        string     `json:"task"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type tasks []Task

func (toDos *tasks) add(task string) {
	todo := Task{
		Task:        strings.TrimSpace(task),
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*toDos = append(*toDos, todo)
}

func (toDos *tasks) remove(index int) {
	*toDos = append((*toDos)[:index], (*toDos)[index+1:]...)
}

func (toDos *tasks) update(index int, task string) {
	(*toDos)[index].Task = task
}

func (toDos *tasks) markCompleted(index int) {
	(*toDos)[index].Completed = true
	completedTime := time.Now()
	(*toDos)[index].CompletedAt = &completedTime
}

func (toDos *tasks) clearList() {
	*toDos = tasks{}
}

func (toDos *tasks) list() []Task {
	fmt.Println("To-Do List!")
	fmt.Println("-----------------------------------------------------------")
	if len(*toDos) == 0 {
		fmt.Println("You have no tasks!")
	}

	for index, task := range *toDos {
		completedStatus := "[ ]"
		if task.Completed {
			completedStatus = "[X]"
		}

		createdTime := "Created at: " + task.CreatedAt.Format(time.RFC1123)

		completedTime := "Not Completed!"

		if task.CompletedAt != nil {
			completedTime = "Completed at: " + task.CompletedAt.Format(time.RFC1123)
		}

		fmt.Print("Task: ", index+1, "\n", task.Task, "\n", completedStatus, "\n", createdTime, "\n", completedTime, "\n")
		fmt.Println("-----------------------------------------------------------")
	}
	return *toDos
}

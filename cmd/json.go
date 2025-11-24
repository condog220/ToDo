package cmd

import (
	"encoding/json"
	"os"
)

func SavetoJson(toDos tasks) error {
	data, err := json.MarshalIndent(toDos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

func LoadfromJson() (tasks, error) {
	var toDos tasks
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return toDos, err
	}

	err = json.Unmarshal(data, &toDos)
	return toDos, err
}

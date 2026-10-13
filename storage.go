package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func writeTasks(tasks []Task, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
		return err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("Ошибка: %v", err)
		}
	}()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Printf("Ошибка при сохранении в json: %v", err)
		return err
	}

	return nil
}

func readTasks(fileName string) ([]Task, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil
		}
		fmt.Printf("Ошибка: %v", err)
		return nil, err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("Ошибка: %v", err)
		}
	}()

	var tasks []Task

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		if err == io.EOF {
			return []Task{}, nil
		}
		fmt.Printf("Ошибка: %v", err)
		return nil, err
	}

	return tasks, nil
}

func getTasksPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
		return "", err
	}
	dir := filepath.Join(homeDir, ".gotodo")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Printf("Ошибка при создании файла: %v", err)
		return "", err
	}
	path := filepath.Join(dir, "tasks.json")

	return path, nil
}

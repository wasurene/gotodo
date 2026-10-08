package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	tasks := []Task{}
	maxId := 0

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Введите команду")
			os.Exit(1)
		}
		for i := range tasks {
			if tasks[i].ID > maxId {
				maxId = tasks[i].ID
			}
		}

		newTask := Task{
			ID:       maxId + 1,
			Text:     os.Args[2],
			Done:     false,
			Priority: "",
			Due:      "",
		}

		tasks = append(tasks, newTask)

	case "list":
		fmt.Println("ID", "TEXT", "DONE", "PRIORITY", "DUE")
		for _, task := range tasks {
			if task.Done == true {
				fmt.Println("[X]", task.ID, "|", task.Text, "|", task.Priority, "|", task.Due)
			} else {
				fmt.Println("[]", task.ID, "|", task.Text, "|", task.Priority, "|", task.Due)
			}
		}
	case "done":
		completed := false
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Ошибка: %v", err)
		}
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Done = true
				completed = true
			}
		}
		if completed {
			fmt.Printf("Задача %v выполнена", id)
		} else {
			fmt.Printf("Задача %v не найдена", id)
		}

	case "rm":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Ошибка: %v", err)
		}
		tasks = slices.DeleteFunc(tasks, func(t Task) bool {
			return t.ID == id
		})

	case "clear":
		tasks = []Task{}

	default:
		fmt.Println("Введите команду: add, list, done, rm, clear")
	}

}

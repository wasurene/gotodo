package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {
	path, err := getTasksPath()
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
	}

	tasks, err := readTasks(path)
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
	}

	switch os.Args[1] {
	case "add":
		maxId := 0
		if len(os.Args) < 3 {
			fmt.Println("Введите команду")
			os.Exit(1)
		}

		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		addPriority := addCmd.String("priority", "низкий", "назначение приоритета")
		addDue := addCmd.String("due", "-", "назначить дедлайн")
		addCmd.Parse(os.Args[3:])

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
		newTask.Priority = *addPriority
		newTask.Due = *addDue

		tasks = append(tasks, newTask)

	case "list":
		green := color.GreenString("[X] ")
		red := color.RedString("[] ")

		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		listFilter := listCmd.String("filter", "", "фильтр выполнено/не выполнено")
		listCmd.Parse(os.Args[2:])

		switch strings.ToLower(*listFilter) {
		case "done":
			fmt.Println("   ID", "TEXT", "PRIORITY", "DUE")
			for _, task := range tasks {
				if task.Done == true {
					fmt.Print(green, task.ID, "|", task.Text, "|", task.Priority, "|", task.Due, "\n")
				}
			}
		case "active":
			fmt.Println("   ID", "TEXT", "PRIORITY", "DUE")
			for _, task := range tasks {
				if task.Done == false {
					fmt.Print(red, task.ID, "|", task.Text, "|", task.Priority, "|", task.Due, "\n")
				}
			}
		default:
			fmt.Println("   ID", "TEXT", "PRIORITY", "DUE")
			for _, task := range tasks {
				if task.Done == true {
					fmt.Print(green, task.ID, "|", task.Text, "|", task.Priority, "|", task.Due, "\n")
				} else {
					fmt.Print(red, task.ID, "|", task.Text, "|", task.Priority, "|", task.Due, "\n")
				}
			}
		}
		return
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
		completed := false
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Ошибка: %v", err)
		}
		tasks = slices.DeleteFunc(tasks, func(t Task) bool {
			completed = true
			return t.ID == id
		})
		if completed {
			fmt.Printf("Задача %d удалена", id)
		} else {
			fmt.Printf("Задача %v не найдена", id)
		}

	case "clear":
		tasks = []Task{}

	default:
		fmt.Println("Введите команду: add, list, done, rm, clear")
	}

	err = writeTasks(tasks, path)
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
	}
}

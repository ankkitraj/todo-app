package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	name    string
	checked bool
}

var tasks []Task

func main() {
	fmt.Println("Welcome to your To-Do List!")
	menu()
}

func menu() {
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Exit")

		reader := bufio.NewReader(os.Stdin)
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			addTask()
		case "2":
			listTasks()
		case "3":
			completeTask()
		case "4":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option, please choose again")
		}
	}
}

func addTask() {
	fmt.Println("Enter task name:")
	reader := bufio.NewReader(os.Stdin)
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)
	tasks = append(tasks, Task{name: taskName})
	fmt.Println("Added:", taskName)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks to show!")
		return
	}
	for i, task := range tasks {
		status := "pending"
		if task.checked {
			status = "completed"
		}
		fmt.Printf("%d. %s - %s\n", i+1, task.name, status)
	}
}

func completeTask() {
	fmt.Println("Enter task number to mark as completed:")
	var num int
	fmt.Scan(&num)
	if num-1 < 0 || num-1 >= len(tasks) {
		fmt.Println("Invalid task number!")
		return
	}
	tasks[num-1].checked = true
	fmt.Println("Task completed:", tasks[num-1].name)
}


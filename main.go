package main

import (
	"fmt"
	"os"
)

// Task Struct
type Task struct {
	ID        int
	Title     string
	Completed bool
}

// Global tasks slice
var tasks []Task

// Helper function to split lines from file
func splitLines(s string) []string {
	var lines []string
	start := 0
	for i, c := range s {
		if c == '\n' {
			lines = append(lines, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}

// Add a task
func addTask() {
	var title string
	fmt.Print("Enter task title: ")
	fmt.Scanln(&title)
	id := len(tasks) + 1
	tasks = append(tasks, Task{ID: id, Title: title, Completed: false})
	fmt.Println("Task added!")
}

// View all tasks
func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, t := range tasks {
		status := "Not Done"
		if t.Completed {
			status = "Done"
		}
		fmt.Printf("%d. %s [%s]\n", t.ID, t.Title, status)
	}
}

// Mark a task as completed
func completeTask() {
	var id int
	fmt.Print("Enter task ID to complete: ")
	fmt.Scan(&id)

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Completed = true
			fmt.Println("Task marked as completed.")
			return
		}
	}
	fmt.Println("Task not found.")
}

// Delete a task
func deleteTask() {
	var id int
	fmt.Print("Enter task ID to delete:")
	fmt.Scan(&id)

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted.")
			return
		}
	}
	fmt.Println("Task not found.")
}

// Save tasks to file
func saveTasks() {
	file, err := os.Create("tasks.txt")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	defer file.Close()

	for _, task := range tasks {
		line := fmt.Sprintf("%d|%s|%t\n", task.ID, task.Title, task.Completed)
		file.WriteString(line)
	}
}

// Load tasks from file
func loadTasks() {
	data, err :=
		os.ReadFile("tasks.txt")
	if err != nil {
		return
	}
	lines := string(data)
	for _, line := range splitLines(lines) {
		if line == "" {
			continue
		}

		var t Task
		fmt.Sscanf(line, "%d|%s|%t", &t.ID, &t.Title, &t.Completed)
		tasks = append(tasks, t)
	}
}

// Main function
func main() {
	loadTasks()

	for {
		fmt.Println("\n--- TO DO APP MENU ---")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task Completed")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addTask()
		case 2:
			viewTasks()
		case 3:
			completeTask()
		case 4:
			deleteTask()
		case 5:
			saveTasks()
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

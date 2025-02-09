package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Title    string `json:"title"`
	Done     bool   `json:"done"`
	Category string `json:"category"`
	Priority string `json:"priority"`
}

var todos []Todo

const saveFile = "todos.json"

func main() {
	loadTodos()
	fmt.Println("🎯 Welcome to the Interactive To-Do List! 🎯")
	for {
		showMenu()
		choice := getInput("Choose an option:")
		handleChoice(choice)
	}
}

func showMenu() {
	fmt.Println("\n📋 Menu:")
	fmt.Println("1. List All To-Dos")
	fmt.Println("2. Add a To-Do")
	fmt.Println("3. Mark a To-Do as Done")
	fmt.Println("4. Delete a To-Do")
	fmt.Println("5. Filter To-Dos by Category")
	fmt.Println("6. List Only Pending or Completed To-Dos")
	fmt.Println("7. Search To-Dos by Title")
	fmt.Println("8. Save and Exit")
}

func handleChoice(choice string) {
	switch choice {
	case "1":
		listTodos(todos)
	case "2":
		addTodo()
	case "3":
		markTodoDone()
	case "4":
		deleteTodo()
	case "5":
		filterByCategory()
	case "6":
		listPendingOrCompleted()
	case "7":
		searchTodos()
	case "8":
		saveTodos()
		fmt.Println("✅ To-Dos saved successfully! Goodbye!")
		os.Exit(0)
	default:
		fmt.Println("❌ Invalid choice. Please try again.")
	}
}

func listTodos(todos []Todo) {
	if len(todos) == 0 {
		fmt.Println("📭 No to-dos found!")
		return
	}
	fmt.Println("\n📌 Your To-Do List:")
	for i, todo := range todos {
		status := "Not Done"
		if todo.Done {
			status = "Done"
		}
		fmt.Printf("%d. [%s] %s (Category: %s, Priority: %s)\n", i+1, status, todo.Title, todo.Category, todo.Priority)
	}
}

func addTodo() {
	title := getInput("Enter the title of the to-do:")
	category := getInput("Enter the category (e.g., Work, Personal):")
	priority := getPriorityInput()
	todo := Todo{
		Title:    title,
		Category: category,
		Priority: priority,
		Done:     false,
	}
	todos = append(todos, todo)
	fmt.Println("✅ To-Do added successfully!")
}

func markTodoDone() {
	id := getTodoID("Enter the ID of the to-do to mark as done:")
	if id < 0 || id >= len(todos) {
		fmt.Println("❌ Invalid ID.")
		return
	}
	todos[id].Done = true
	fmt.Println("✅ To-Do marked as done!")
}

func deleteTodo() {
	id := getTodoID("Enter the ID of the to-do to delete:")
	if id < 0 || id >= len(todos) {
		fmt.Println("❌ Invalid ID.")
		return
	}
	todos = append(todos[:id], todos[id+1:]...)
	fmt.Println("🗑️ To-Do deleted successfully!")
}

func filterByCategory() {
	category := getInput("Enter the category to filter by:")
	var filtered []Todo
	for _, todo := range todos {
		if strings.EqualFold(todo.Category, category) {
			filtered = append(filtered, todo)
		}
	}
	listTodos(filtered)
}

func listPendingOrCompleted() {
	status := getInput("Enter 'pending' to list only pending to-dos or 'completed' for completed ones:")
	var filtered []Todo
	for _, todo := range todos {
		if (status == "pending" && !todo.Done) || (status == "completed" && todo.Done) {
			filtered = append(filtered, todo)
		}
	}
	listTodos(filtered)
}

func searchTodos() {
	keyword := getInput("Enter a keyword to search:")
	var results []Todo
	for _, todo := range todos {
		if strings.Contains(strings.ToLower(todo.Title), strings.ToLower(keyword)) {
			results = append(results, todo)
		}
	}
	listTodos(results)
}

func saveTodos() {
	file, err := os.Create(saveFile)
	if err != nil {
		fmt.Println("❌ Error saving to-dos:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(todos); err != nil {
		fmt.Println("❌ Error encoding to-dos:", err)
	}
}

func loadTodos() {
	file, err := os.Open(saveFile)
	if err != nil {
		fmt.Println("📂 No existing to-dos found.")
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&todos); err != nil {
		fmt.Println("❌ Error decoding to-dos:", err)
	}
}

func getInput(prompt string) string {
	fmt.Print(prompt + " ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func getTodoID(prompt string) int {
	idStr := getInput(prompt)
	var id int
	fmt.Sscanf(idStr, "%d", &id)
	return id - 1
}

func getPriorityInput() string {
	for {
		priority := getInput("Enter the priority (High, Medium, Low):")
		switch strings.ToLower(priority) {
		case "high", "medium", "low":
			return (priority)
		default:
			fmt.Println("❌ Invalid priority. Please enter 'High', 'Medium', or 'Low'.")
		}
	}
}

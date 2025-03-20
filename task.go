package main

// Task represents a unit of work with a unique ID
// Add more fields
type Task struct {
	ID int
}

// getTasks generates a slice of tasks with sequential IDs
// This function should be replaced with one that generates tasks dynamically or fetches them from a database, API, etc.
func getTasks(numTasks int) []Task {
	tasks := make([]Task, 0, numTasks)
	for i := 1; i <= numTasks; i++ {
		tasks = append(tasks, Task{ID: i})
	}

	return tasks
}

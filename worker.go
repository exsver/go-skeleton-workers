package main

import (
	"sync"
	"time"
)

// worker processes tasks from the tasks channel until it receives a stop signal or until no more tasks are left in the channel
func worker(id int, tasks <-chan Task, wg *sync.WaitGroup, stopChan <-chan struct{}) {
	defer func() {
		Log.Info.Printf("Worker %d: Stopped\n", id)

		wg.Done()
	}()

	Log.Info.Printf("Worker %d: Started\n", id)

	for task := range tasks {
		select {
		case <-stopChan:
			Log.Info.Printf("Worker %d: Received stop signal\n", id)
			return
		default:
			Log.Info.Printf("Worker %d: Processing task %d\n", id, task.ID)
			// Processing simulation
			time.Sleep(time.Second * 3)
			// Add app logic here
		}
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

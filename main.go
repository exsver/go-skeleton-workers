package main

import (
	"context"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

func main() {
	// Parse flags
	flags := parseFlags()

	// Set log-level
	setLogger(*flags.LogLevel)

	// Read config
	config, err := getConfiguration(flags)
	if err != nil {
		Log.Error.Fatal(err)
	}

	// Set GOMAXPROCS
	Log.Info.Printf("Setting GOMAXPROCS=%d", config.MaxProcs)
	runtime.GOMAXPROCS(config.MaxProcs)

	// Get tasks
	// This function should be replaced with one that generates tasks dynamically or fetches them from a database, API, etc.
	tasks := getTasks(20)

	// Create and fill taskChan
	taskChan := make(chan Task, len(tasks))
	for _, task := range tasks {
		taskChan <- task
	}

	close(taskChan)

	// Capture termination signals to gracefully shutdown workers
	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	defer stop()

	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= minInt(config.Workers, len(tasks)); i++ {
		wg.Add(1)
		go worker(ctx, i, taskChan, &wg)
	}

	// Wait for all workers to complete
	wg.Wait()

	Log.Info.Println("Finished.")
}

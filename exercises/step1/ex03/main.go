package main

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/gnir-work/learning-go/exercises/step1/ex03/pool"
)

type httpRequestJob struct {
	url string
}

func (request *httpRequestJob) Execute() {
	slog.Info("Sending HTTP request to", "url", request.url)
	time.Sleep(1 * time.Second)
	slog.Info("Finished HTTP request to", "url", request.url)
}

func main() {
	slog.Info("Starting workerPool")
	workerPool := pool.NewPool(5)

	go workerPool.Start()
	for i := 0; i < 10; i += 1 {
		slog.Info("Sending event to worker pool")
		workerPool.Submit(&httpRequestJob{
			url: fmt.Sprintf("https://www.google.com/%v", i),
		})
	}
	slog.Info("Closing workerPool")
	workerPool.Close()
	slog.Info("Waiting for workerPool to finish jobs")
	workerPool.Wait()
	slog.Info("Exiting main")
}

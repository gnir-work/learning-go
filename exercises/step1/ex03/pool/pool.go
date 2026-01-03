package pool

import (
	"log/slog"
	"sync"
)

type Job interface {
	Execute()
}

type Pool struct {
	jobs    chan Job
	workers uint32
	wg      sync.WaitGroup
}

func worker(jobs chan Job, ID uint32) {
	logger := slog.With("worker_id", ID)
	for job := range jobs {
		logger.Info("Executing job")
		job.Execute()
		logger.Info("Finished executing job")
	}
}

func (pool *Pool) queueScheduler(bufferedJobs chan Job) {
	var q []Job
	for {
		var next Job
		var out chan Job
		if len(q) > 0 {
			next = q[0]
			out = bufferedJobs
		}
		select {
		case job, ok := <-pool.jobs:
			if !ok {
				slog.Info("Gracefully closing scheduler as jobs channel was closed")
				if len(q) > 0 {
					slog.Info("Draining queue")
					for _, job := range q {
						bufferedJobs <- job
					}
				}
				slog.Info("Closing bufferedJobs channel")
				close(bufferedJobs)
				return
			}
			slog.Info("Scheduling job")
			q = append(q, job)
		case out <- next:
			q = q[1:]
		}
	}
}

func NewPool(workers uint32) *Pool {
	return &Pool{
		jobs:    make(chan Job, 1024),
		wg:      sync.WaitGroup{},
		workers: workers,
	}
}

func (pool *Pool) Start() {
	bufferedJobs := make(chan Job, pool.workers)
	for i := uint32(0); i < pool.workers; i++ {
		pool.wg.Go(func() {
			worker(bufferedJobs, i)
		})
	}
	pool.wg.Go(func() {
		pool.queueScheduler(bufferedJobs)
	})
}

func (pool *Pool) Submit(job Job) {
	pool.jobs <- job
}

func (pool *Pool) Wait() {
	pool.wg.Wait()
}

func (pool *Pool) Close() {
	close(pool.jobs)
}

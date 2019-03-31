package go_pool

import (
	"log"
	"time"
)

// Worker is the actual executor who runs the tasks,
// it starts a goroutine that accepts tasks and
// performs function calls.
type Worker struct {
	// pool who owns this worker.
	pool *Pool

	// task is a job should be done.
	task chan func()

	// recycleTime will be update when putting a worker back into queue.
	recycleTime time.Time
}

// run starts a goroutine to repeat the process
// that performs the function calls.
func (w *Worker) run() {
	w.pool.incRunning()
	go func() {
		defer func() {
			if p := recover(); p != nil {
				w.pool.decRunning()
				if w.pool.PanicHandler != nil {
					w.pool.PanicHandler(p)
				} else {
					log.Printf("worker exits from a panic: %v", p)
				}
			}
		}()

		for f := range w.task {
			if f == nil {
				w.pool.decRunning()
				w.pool.workerCache.Put(w)
				return
			}
			f()
			w.pool.revertWorker(w)
		}
	}()
}

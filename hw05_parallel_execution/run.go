package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	wg := sync.WaitGroup{}
	maxErrors := int32(m)
	ch := producer(tasks)

	// or replace into the consumer and wg.Add(1) ???
	wg.Add(n)
	// as for me not pretty pass by ref the counter, but how else (global)???
	var errCounter int32
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			consumer(ch, &errCounter, maxErrors)
		}()
	}
	wg.Wait()

	if errCounter >= maxErrors {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func producer(tasks []Task) chan Task {
	ch := make(chan Task, len(tasks))

	// maybe goroutine is unnecessary ???
	go func() {
		for _, task := range tasks {
			ch <- task
		}
		defer close(ch)
	}()
	return ch
}

func consumer(ch chan Task, errCounter *int32, maxErrors int32) {
	for task := range ch {
		if err := task(); err != nil {
			// avoid race conditions
			if atomic.AddInt32(errCounter, 1) > maxErrors {
				return
			}
		}
	}
}

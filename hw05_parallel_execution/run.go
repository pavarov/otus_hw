package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	wg := sync.WaitGroup{}
	maxErrors := int32(m)
	ch := producer(tasks)

	wg.Add(n)
	var errCounter int32
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range ch {
				if err := task(); err != nil {
					// avoid race conditions
					if atomic.AddInt32(&errCounter, 1) > maxErrors {
						return
					}
				}
			}
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
	defer close(ch)
	for _, task := range tasks {
		ch <- task
	}
	return ch
}

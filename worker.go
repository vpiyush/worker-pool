package main

import (
	"log"
	"time"
)

type Worker struct {
	id          int
	workQueue   chan Work
	workerQueue chan chan Work
	quitChan    chan bool
}

func NewWorker(i int, wq chan chan Work) *Worker {
	return &Worker{
		id:          i,
		workQueue:   make(chan Work),
		workerQueue: wq,
		quitChan:    make(chan bool),
	}
}
func (w Worker) Start() {
	// add our serlves into worker queue
	w.workerQueue <- w.workQueue
	go func() {
		for {
			select {
			case work := <-w.workQueue:
				log.Println("Work received")
				w.workQueue <- work
				time.Sleep(work.Delay)
			case <-w.quitChan:
				log.Println("stopping the worker ", w.id)
			}
		}
	}()
}

func (w Worker) Stop() {
	w.quitChan <- true
}

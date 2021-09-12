package main

import "log"

func StartDispatcher(numWorker int) {

	// initialize the woker pool
	workerQueue := make(chan chan Work, numWorker)

	// now initialize all workers inside the pool
	for i := 0; i < numWorker; i++ {
		log.Println("Starting Worker ", i+1)
		worker := NewWorker(i+1, workerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				log.Println("Received work")
				// start another go routine to push this work
				// so that it doesn't wait if workers aren't available
				go func() {
					worker := <-workerQueue
					log.Println("Dispatching work request ")
					worker <- work
				}()
			}
		}
	}()

}

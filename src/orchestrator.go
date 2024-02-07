package main

import "time"

type DatastoreAccessOrchestrator struct {
	UpdatePersonsStagingQueue    chan UpdatePersonQueueItem
	UpdatePersonsProcessingQueue chan UpdatePersonQueueItem
	GetPersonsStagingQueue       chan GetPersonQueueItem
	GetPersonsProcessingQueue    chan GetPersonQueueItem
	Datastore                    *PersonDatastore
}

func (d *DatastoreAccessOrchestrator) Start() {
	for {
		if len(d.GetPersonsStagingQueue) > 0 {
			// Try to clear the update processing queue and then push onto the get processing queue
			if len(d.UpdatePersonsProcessingQueue) > 0 {
				// Wait until this queue is empty
				time.Sleep(10 * time.Millisecond)
			} else {
				// Get the person from the datastore
				d.GetPersonsProcessingQueue <- <-d.GetPersonsStagingQueue
			}
		} else if len(d.UpdatePersonsStagingQueue) > 0 {
			// Try to clear the get processing queue and then push onto the update processing queue
			if len(d.GetPersonsProcessingQueue) > 0 {
				// Wait until this queue is empty
				time.Sleep(10 * time.Millisecond)
			} else {
				// Get the person from the datastore
				d.UpdatePersonsProcessingQueue <- <-d.UpdatePersonsStagingQueue
			}
		}
	}
}

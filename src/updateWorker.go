package main

type UpdateWorker struct {
	UpdatePersonsProcessingQueue chan UpdatePersonQueueItem
	PersonDatastore              *PersonDatastore
	// Later we can set an is working flag so that we can easily tell if any update is running before we perform a read
	// or we make a updateWorkersManager and readWorkersManager with the ability to grab a lock on the datastore.
}

type UpdatePersonQueueItem struct {
	UpdatePersonEntity UpdatePersonEntity
	ReturnChannel      chan PersonLog
}

func (u *UpdateWorker) Start() {
	for {
		if len(u.UpdatePersonsProcessingQueue) > 0 {
			updatePersonItem := <-u.UpdatePersonsProcessingQueue
			log := u.PersonDatastore.Update(updatePersonItem.UpdatePersonEntity)
			updatePersonItem.ReturnChannel <- log
		}
	}
}

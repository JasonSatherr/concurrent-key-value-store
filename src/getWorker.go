package main

type GetWorker struct {
	GetPersonsProcessingQueue chan GetPersonQueueItem
	PersonDatastore           *PersonDatastore
}

type GetPersonQueueItem struct {
	ID            int
	ReturnChannel chan PersonLog
}

func (g *GetWorker) Start() {
	for {
		if len(g.GetPersonsProcessingQueue) > 0 {
			// Get the person from the datastore
			// and send it to the getPersonsQueue
			getPersonItem := <-g.GetPersonsProcessingQueue
			personLog := g.PersonDatastore.Get(getPersonItem.ID)
			getPersonItem.ReturnChannel <- personLog
		}
	}
}

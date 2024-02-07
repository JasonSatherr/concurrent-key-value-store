package main

import (
	"log"
	"net/http"
)

func main() {
	personDatastore := GetNewPersonDatastore()
	getPersonStagingQueue := make(chan GetPersonQueueItem, 200)
	getPersonProcessingQueue := make(chan GetPersonQueueItem, 200)
	updatePersonStagingQueue := make(chan UpdatePersonQueueItem, 100)
	updatePersonProcessingQueue := make(chan UpdatePersonQueueItem, 100)

	getPersonHandler := GetPersonHandler{GetPersonsStagingQueue: getPersonStagingQueue}
	updatePersonHandler := UpdatePersonHandler{UpdatePersonsStagingQueue: updatePersonStagingQueue}

	getWorker := GetWorker{GetPersonsProcessingQueue: getPersonProcessingQueue, PersonDatastore: personDatastore}
	updateWorker := UpdateWorker{UpdatePersonsProcessingQueue: updatePersonProcessingQueue, PersonDatastore: personDatastore}

	go getWorker.Start()
	go updateWorker.Start()

	orchestrator := DatastoreAccessOrchestrator{
		UpdatePersonsStagingQueue:    updatePersonStagingQueue,
		UpdatePersonsProcessingQueue: updatePersonProcessingQueue,
		GetPersonsStagingQueue:       getPersonStagingQueue,
		GetPersonsProcessingQueue:    getPersonProcessingQueue,
		Datastore:                    personDatastore,
	}

	go orchestrator.Start()

	mux := http.NewServeMux()
	mux.Handle("/person/update", updatePersonHandler)
	mux.Handle("/person/get", getPersonHandler)

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

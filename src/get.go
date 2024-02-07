package main

import (
	"encoding/json"
	"net/http"
)

type getPersonRequest struct {
	ID int `json:"id"`
}

type GetPersonHandler struct {
	GetPersonsStagingQueue chan GetPersonQueueItem
}

func (g GetPersonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var getPersonRequest = getPersonRequest{}

	err := json.NewDecoder(r.Body).Decode(&getPersonRequest)
	if err != nil {
		http.Error(w, error.Error(err), http.StatusBadRequest)
		return
	}

	getPersonQueueItem := GetPersonQueueItem{
		ID:            getPersonRequest.ID,
		ReturnChannel: make(chan PersonLog),
	}

	g.GetPersonsStagingQueue <- getPersonQueueItem

	json.NewEncoder(w).Encode(<-getPersonQueueItem.ReturnChannel)
}

package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type UpdatePersonEntity struct {
	ID        int
	Name      string
	UpdatedAt time.Time
}

type UpdatePersonHandler struct {
	UpdatePersonsStagingQueue chan UpdatePersonQueueItem
}

func (u UpdatePersonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var p = UpdatePersonEntity{}

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, error.Error(err), http.StatusBadRequest)
		return
	}

	p.UpdatedAt = time.Now()

	updatePersonQueueItem := UpdatePersonQueueItem{
		UpdatePersonEntity: p,
		ReturnChannel:      make(chan PersonLog),
	}

	u.UpdatePersonsStagingQueue <- updatePersonQueueItem
	json.NewEncoder(w).Encode(<-updatePersonQueueItem.ReturnChannel)
}

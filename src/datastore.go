package main

import (
	"sync"
	"time"
)

type PersonLog struct {
	Name      string
	UpdatedAt time.Time
}

type PersonDatastore struct {
	mu        sync.Mutex
	datastore map[int]PersonLog
}

func GetNewPersonDatastore() *PersonDatastore {
	return &PersonDatastore{datastore: make(map[int]PersonLog)}
}

func (pd *PersonDatastore) Update(updateEntity UpdatePersonEntity) PersonLog {
	pd.mu.Lock()
	defer pd.mu.Unlock()
	if log, found := (*pd).datastore[updateEntity.ID]; !found || log.UpdatedAt.Before(updateEntity.UpdatedAt) {
		(*pd).datastore[updateEntity.ID] = PersonLog{Name: updateEntity.Name, UpdatedAt: updateEntity.UpdatedAt}
	}

	return (*pd).datastore[updateEntity.ID]
}

func (pd *PersonDatastore) Get(id int) PersonLog {
	return (*pd).datastore[id]
}

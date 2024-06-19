package repository

import (
	"errors"
	"level_2/develop/dev11/model"
	"time"
)

type EventRepository interface {
	CreateEvent(event model.Event) int
	UpdateEvent(id int, event model.Event) error
	DeleteEvent(id int) error
	GetEvents(userID int, startDate, endDate time.Time) []model.Event
}

type EventRepositoryImpl struct {
	storage *Storage
}

func NewEventRepositoryImpl(storage *Storage) *EventRepositoryImpl {
	return &EventRepositoryImpl{storage: storage}
}

func (r *EventRepositoryImpl) CreateEvent(event model.Event) int {
	r.storage.id++
	r.storage.events[r.storage.id] = event
	return r.storage.id
}

func (r *EventRepositoryImpl) UpdateEvent(id int, event model.Event) error {
	_, exist := r.storage.events[id]
	if !exist {
		return errors.New("this event is not existing")
	}
	r.storage.events[id] = event
	return nil
}

func (r *EventRepositoryImpl) DeleteEvent(id int) error {
	_, exist := r.storage.events[id]
	if !exist {
		return errors.New("this event is not existing")
	}
	delete(r.storage.events, id)
	return nil
}

func (r *EventRepositoryImpl) GetEvents(userID int, startDate, endDate time.Time) []model.Event {
	var filteredEvents []model.Event
	for _, event := range r.storage.events {
		if event.UserID == userID && (event.Date.After(startDate) || event.Date.Equal(startDate)) && event.Date.Before(endDate) {
			filteredEvents = append(filteredEvents, event)
		}
	}
	return filteredEvents
}

type Repository struct {
	EventRepository
}

func NewRepository(storage *Storage) *Repository {
	return &Repository{NewEventRepositoryImpl(storage)}
}

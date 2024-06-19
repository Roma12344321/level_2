package service

import (
	"level_2/develop/dev11/model"
	"level_2/develop/dev11/repository"
	"time"
)

type EventService interface {
	CreateEvent(event model.Event) int
	UpdateEvent(id int, event model.Event) error
	DeleteEvent(id int) error
	GetEvents(userID int, startDate, endDate time.Time) []model.Event
}

type EventServiceImpl struct {
	repo *repository.Repository
}

func NewEventServiceImpl(repo *repository.Repository) *EventServiceImpl {
	return &EventServiceImpl{repo: repo}
}

func (s *EventServiceImpl) CreateEvent(event model.Event) int {
	return s.repo.EventRepository.CreateEvent(event)
}

func (s *EventServiceImpl) UpdateEvent(id int, event model.Event) error {
	return s.repo.EventRepository.UpdateEvent(id, event)
}

func (s *EventServiceImpl) DeleteEvent(id int) error {
	return s.repo.EventRepository.DeleteEvent(id)
}

func (s *EventServiceImpl) GetEvents(userID int, startDate, endDate time.Time) []model.Event {
	return s.repo.EventRepository.GetEvents(userID, startDate, endDate)
}

type Service struct {
	EventService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{EventService: NewEventServiceImpl(repo)}
}

package repository

import "level_2/develop/dev11/model"

type Storage struct {
	events map[int]model.Event
	id     int
}

func NewStorage() *Storage {
	return &Storage{events: make(map[int]model.Event), id: 0}
}

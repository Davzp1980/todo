package service

import "todo/repository"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

// Сервисы обращаются к базе данных
func NewService(repos *repository.Repository) *Service {
	return &Service{}
}

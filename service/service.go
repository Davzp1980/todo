package service

import "github.com/Davzp1980/todo-app/repository"

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

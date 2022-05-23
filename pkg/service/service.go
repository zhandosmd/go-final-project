package service

import (
	"github.com/zhandosmd/golang-todo/pkg/repository"
)

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}

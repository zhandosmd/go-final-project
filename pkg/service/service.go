package service

import (
	lofo "github.com/zhandosmd/go-final-project"
	"github.com/zhandosmd/go-final-project/pkg/repository"
)

type Authorization interface {
	CreateUser(User lofo.User) (int, error)
	GenerateToken(Username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list lofo.Lofo) (int, error)
	GetAll(userId int) ([]lofo.Lofo, error)
	GetById(userId, listId int) (lofo.Lofo, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input lofo.UpdateListInput) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.Lofo),
	}
}

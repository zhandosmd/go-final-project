package service

import (
	todo "github.com/zhandosmd/go-final-project"
	"github.com/zhandosmd/go-final-project/pkg/repository"
)

type Authorization interface {
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	Update(userId, listId int, input todo.UpdateListInput) error
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoList: NewTodoListService(repos.TodoList),
	}
}

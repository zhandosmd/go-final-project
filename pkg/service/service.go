package service

import (
	todo "github.com/zhandosmd/golang-todo"
	"github.com/zhandosmd/golang-todo/pkg/repository"
)

type Authorization interface {
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	Update(userId, listId int, input todo.UpdateListInput) error
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

package repository

import (
	todo "github.com/zhandosmd/golang-todo"
)

type Authorization interface {
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	Update(userId, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository() *Repository {
	return &Repository{
		TodoList: NewTodoListPostgres(db),
	}
}

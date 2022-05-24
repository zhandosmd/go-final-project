package repository

import (
	"github.com/jmoiron/sqlx"
	lofo "github.com/zhandosmd/go-final-project"
)

type Authorization interface {
	CreateUser(User lofo.User) (int, error)
	GetUser(username, password string) (lofo.User, error)
}

type Lofo interface {
	Create(userId int, list lofo.Lofo) (int, error)
	GetAll(userId int) ([]lofo.Lofo, error)
	GetById(userId, listId int) (lofo.Lofo, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input lofo.UpdateListInput) error
}

type Repository struct {
	Authorization
	Lofo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Lofo:          NewTodoListPostgres(db),
	}
}

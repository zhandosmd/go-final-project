package service

import (
	lofo "github.com/zhandosmd/go-final-project"
	"github.com/zhandosmd/go-final-project/pkg/repository"
)

type LofoService struct {
	repo repository.Lofo
}

func NewTodoListService(repo repository.Lofo) *LofoService {
	return &LofoService{repo: repo}
}

func (s *LofoService) Create(userId int, list lofo.Lofo) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *LofoService) GetAll(userId int) ([]lofo.Lofo, error) {
	return s.repo.GetAll(userId)
}

func (s *LofoService) GetById(userId, listId int) (lofo.Lofo, error) {
	return s.repo.GetById(userId, listId)
}

func (s *LofoService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *LofoService) Update(userId, listId int, input lofo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}

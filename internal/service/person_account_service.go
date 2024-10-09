package service

import (
	"context"
	"gomicrobase/internal/model"
	"gomicrobase/internal/repository"
)

type PersonAccountService struct {
	repo repository.PersonAccountRepository
}

func NewPersonAccountService(r repository.PersonAccountRepository) *PersonAccountService {
	return &PersonAccountService{repo: r}
}

func (s *PersonAccountService) CreatePerson(ctx context.Context, person *model.Person) (*model.Person, error) {
	err := s.repo.CreatePerson(person)
	return person, err
}

func (s *PersonAccountService) GetPerson(ctx context.Context, id int32) (*model.Person, error) {
	return s.repo.GetPerson(id)
}

// Similar functions for Account

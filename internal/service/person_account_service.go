package service

import (
	"context"
	"gomicrobase/internal/model"
	"gomicrobase/internal/repository"
)

type PersonAccountServiceInterface interface {
	CreatePerson(ctx context.Context, person *model.Person) (*model.Person, error)
	GetPerson(ctx context.Context, id int32) (*model.Person, error)
	CreateAccount(ctx context.Context, cuenta *model.Account) (*model.Account, error)
	GetAccount(ctx context.Context, id int32) (*model.Account, error)
}

type PersonAccountService struct {
	repo *repository.PersonAccountRepository
}

func NewPersonAccountService(repo repository.PersonAccountRepository) PersonAccountService {
	return PersonAccountService{repo: &repo}
}

func (s *PersonAccountService) CreatePerson(ctx context.Context, person *model.Person) (*model.Person, error) {
	err := s.repo.CreatePerson(person)
	return person, err
}

func (s *PersonAccountService) GetPerson(ctx context.Context, id int32) (*model.Person, error) {
	return s.repo.GetPerson(id)
}

// Similar functions for Account

func (s *PersonAccountService) CreateAccount(ctx context.Context, cuenta *model.Account) (*model.Account, error) {
	err := s.repo.CreateAccount(cuenta)
	return cuenta, err
}

func (s *PersonAccountService) GetAccount(ctx context.Context, id int32) (*model.Account, error) {
	return s.repo.GetAccount(id)
}

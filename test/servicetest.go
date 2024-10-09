package service_test

import (
	"context"
	"gomicrobase/internal/model"
	"gomicrobase/internal/repository"
	"gomicrobase/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePerson(t *testing.T) {
	repo := repository.NewMockPersonAccountRepository()
	svc := service.NewPersonAccountService(repo)

	person := &model.Person{Name: "John Doe", Email: "john@example.com"}
	createdPerson, err := svc.CreatePerson(context.Background(), person)

	assert.NoError(t, err)
	assert.Equal(t, person.Name, createdPerson.Name)
	assert.Equal(t, person.Email, createdPerson.Email)
}

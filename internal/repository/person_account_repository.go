package repository

import (
	"database/sql"
	"gomicrobase/internal/model"
)

type PersonAccountRepository struct {
	db *sql.DB
}

func NewPersonAccountRepository(db *sql.DB) *PersonAccountRepository {
	return &PersonAccountRepository{db: db}
}

func (r *PersonAccountRepository) CreatePerson(person *model.Person) error {
	_, err := r.db.Exec("INSERT INTO persons (name, email) VALUES ($1, $2)", person.Name, person.Email)
	return err
}

func (r *PersonAccountRepository) GetPerson(id int32) (*model.Person, error) {
	row := r.db.QueryRow("SELECT id, name, email FROM persons WHERE id = $1", id)
	person := &model.Person{}
	err := row.Scan(&person.ID, &person.Name, &person.Email)
	return person, err
}

// Similar functions for Account

package repository

import (
	"database/sql"
	"gomicrobase/internal/model"
)

type PersonAccountRepositoryInterface interface {
	CreatePerson(person *model.Person) error
	GetPerson(id int32) (*model.Person, error)
	CreateAccount(cuenta *model.Account) error
	GetAccount(id int32) (*model.Account, error)
}

type PersonAccountRepository struct {
	db *sql.DB
}

func NewPersonAccountRepository(db *sql.DB) PersonAccountRepository {
	return PersonAccountRepository{db: db}
}

func (r *PersonAccountRepository) CreatePerson(person *model.Person) error {
	_, err := r.db.Exec("INSERT INTO personas (nombre, email) VALUES ($1, $2)", person.Name, person.Email)
	return err
}

func (r *PersonAccountRepository) GetPerson(id int32) (*model.Person, error) {
	row := r.db.QueryRow("SELECT id, nombre, email FROM personas WHERE id = $1", id)
	person := &model.Person{}
	err := row.Scan(&person.ID, &person.Name, &person.Email)
	return person, err
}

// Similar functions for Account

func (r *PersonAccountRepository) CreateAccount(cuenta *model.Account) error {
	_, err := r.db.Exec("INSERT INTO cuentas (persona_id, saldo) VALUES ($1, $2)", cuenta.PersonID, cuenta.Balance)
	return err
}

func (r *PersonAccountRepository) GetAccount(id int32) (*model.Account, error) {
	row := r.db.QueryRow("SELECT id, persona_id, saldo FROM cuentas WHERE id = $1", id)
	cuenta := &model.Account{}
	err := row.Scan(&cuenta.ID, &cuenta.PersonID, &cuenta.Balance)
	return cuenta, err
}

package handler

import "github.com/PACOdergod/crud-golang/models"

type Storage interface {
	Create(person *models.Person) error
	Update(id int, person *models.Person) error
	Delete(id int) error
	GetByID(id int) (models.Person, error)
	GetAll() (models.Persons, error)
}

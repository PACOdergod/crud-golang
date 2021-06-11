package storage

import (
	"github.com/PACOdergod/crud-golang/models"
)

type persona = models.Person

func NewMemory() Memory {
	personas := make(map[int]persona)

	return Memory{
		CurrentID: 0,
		Persons:   personas,
	}
}

type Memory struct {
	CurrentID int
	Persons   map[int]persona
}

// Create
func (m *Memory) Create(person *persona) error {
	if person == nil {
		return models.ErrPersonNotNull
	}

	m.CurrentID++
	m.Persons[m.CurrentID] = *person

	return nil
}

// Actualizar
func (m *Memory) Update(id int, person *persona) error {
	if person == nil {
		return models.ErrPersonNotNull
	}

	if _, ok := m.Persons[id]; !ok {
		return models.ErrPersonNotExist
	}

	m.Persons[id] = *person

	return nil
}

// Borrar
func (m *Memory) Delete(id int) error {
	if _, ok := m.Persons[id]; !ok {
		return models.ErrPersonNotExist
	}

	delete(m.Persons, id)

	return nil
}

// Leer por Id
func (m *Memory) GetByID(id int) (persona, error) {
	person, ok := m.Persons[id]

	if !ok {
		return person, models.ErrPersonNotExist
	}

	return person, nil
}

func (m *Memory) GetAll() (models.Persons, error) {
	var result models.Persons

	for _, v := range m.Persons {
		result = append(result, v)
	}

	return result, nil
}

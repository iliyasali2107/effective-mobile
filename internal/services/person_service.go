package services

import (
	"context"
	"effective-mobile/internal/domain/dto"
	"effective-mobile/internal/domain/models"
	"effective-mobile/internal/storage"
	"errors"
)

type PersonSvc struct {
	personStorage PersonStorage
}

type PersonStorage interface {
	AddPerson(ctx context.Context, person models.Person) (int, error)
	DeletePerson(ctx context.Context, personId int) error
	UpdatePerson(ctx context.Context, person models.Person) error

	// TODO: not finished
	GetPersonsByFilter(ctx context.Context) ([]models.Person, error)
}

func NewPersonSvc(storage PersonStorage) *PersonSvc {
	return &PersonSvc{
		personStorage: storage,
	}
}

func (ps *PersonSvc) AddPerson(ctx context.Context, req dto.AddPersonRequest) (int, error) {
	person := addPersonRequestToModel(req)
	id, err := ps.personStorage.AddPerson(ctx, person)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ps *PersonSvc) DeletePerson(ctx context.Context, personId int) error {
	return ps.personStorage.DeletePerson(ctx, personId)
}

func (ps *PersonSvc) UpdatePerson(ctx context.Context, req dto.UpdatePersonRequest) error {
	person := updatePersonRequestToModel(req)
	err := ps.personStorage.UpdatePerson(ctx, person)
	if err != nil {
		if errors.Is(err, storage.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}

	return nil
}

func addPersonRequestToModel(req dto.AddPersonRequest) models.Person {
	person := models.Person{}

	person.Age = req.Age
	person.Gender = req.Gender
	person.Name = req.Name
	person.Surname = req.Surname
	person.Patronymic = req.Patronymic
	person.Nationality = req.Nationality

	return person
}

func updatePersonRequestToModel(req dto.UpdatePersonRequest) models.Person {
	person := models.Person{}
	person.Id = req.Id
	person.Age = req.Age
	person.Gender = req.Gender
	person.Name = req.Name
	person.Surname = req.Surname
	person.Patronymic = req.Patronymic
	person.Nationality = req.Nationality

	return person
}

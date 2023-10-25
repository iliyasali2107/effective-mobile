package storage

import (
	"context"
	"database/sql"
	"effective-mobile/internal/domain/models"
	"fmt"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (ps *PostgresStorage) AddPerson(ctx context.Context, person models.Person) (int, error) {
	query := `INSERT INTO persons(name, surname, patronymic, age, gender, nationality) VALUES($1, $2, $3, $4, $5, $6) RETURNING id;`

	args := []interface{}{person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality}
	row := ps.db.QueryRowContext(ctx, query, args...)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (ps *PostgresStorage) DeletePerson(ctx context.Context, personId int) error {
	query := `DELETE FROM persons WHERE id = $1;`

	_, err := ps.db.ExecContext(ctx, query, personId)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PostgresStorage) UpdatePerson(ctx context.Context, person models.Person) error {
	acc := 1
	query := "UPDATE persons SET "
	args := []interface{}{}
	if person.Name != "" {
		query += fmt.Sprintf("name = $%d, ", acc)
		acc++
		args = append(args, person.Name)
	}
	if person.Surname != "" {
		query += fmt.Sprintf("surname = $%d, ", acc)
		acc++
		args = append(args, person.Surname)
	}
	if person.Patronymic != "" {
		query += fmt.Sprintf("patronymic = $%d, ", acc)
		acc++
		args = append(args, person.Patronymic)
	}
	if person.Age != 0 {
		query += fmt.Sprintf("age = $%d, ", acc)
		acc++
		args = append(args, person.Age)
	}
	if person.Gender != "" {
		query += fmt.Sprintf("gender = $%d, ", acc)
		acc++
		args = append(args, person.Gender)
	}
	if person.Nationality != "" {
		query += fmt.Sprintf("nationality = $%d, ", acc)
		acc++
		args = append(args, person.Nationality)
	}
	if len(args) == 0 {
		return ErrNoUpdateData
	}

	query = query[:len(query)-2]
	query += fmt.Sprintf(" WHERE id = $%d", acc)
	args = append(args, person.Id)

	res, err := ps.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNoRows
	}

	return nil
}

// TODO: not finished
func (ps *PostgresStorage) GetPersonsByFilter(ctx context.Context) ([]models.Person, error) {
	return nil, nil
}

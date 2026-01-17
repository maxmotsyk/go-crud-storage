package stor

import (
	"database/sql"
	"errors"
	"fmt"
	"gocrud/internal/domain"
	"time"

	log "github.com/sirupsen/logrus"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(d *sql.DB) *Storage {
	return &Storage{
		DB: d,
	}
}

func (s *Storage) CreateUser(u *domain.SignUpInput) error {
	registered_time := time.Now()
	_, err := s.DB.Exec(`
        INSERT INTO users (name, last_name, age, email, password, registered_time)
        VALUES ($1, $2, $3, $4, $5, $6)
    `, u.Name, u.LastName, u.Age, u.Email, u.Password, registered_time)

	if err != nil {
		log.Printf("CreateUser DB error: %v", err)
		return err
	}

	log.WithFields(log.Fields{
		"layer":   "storage",
		"action":  "createUser",
		"user_id": u.Email,
	}).Info("Successfully created user")

	return nil
}

func (s *Storage) GetUser(id int64) (*sql.Rows, error) {

	row, err := s.DB.Query("SELECT id, name, last_name, age, email, registered_time FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"layer":  "storage",
		"action": "getUser",
	}).Info(fmt.Sprintf("Successfully retrieved user with id %d", id))

	return row, nil
}

func (s *Storage) UpdateUser(u *domain.UpdateUserInput, id int64) error {
	query := "UPDATE users SET "
	args := []interface{}{}
	argIndex := 1

	// Динамічно додаємо поля, які були передані (не пусті)
	if u.Name != "" {
		query += fmt.Sprintf("name = $%d, ", argIndex)
		args = append(args, u.Name)
		argIndex++
	}

	if u.LastName != "" {
		query += fmt.Sprintf("last_name = $%d, ", argIndex)
		args = append(args, u.LastName)
		argIndex++
	}

	if u.Age != 0 {
		query += fmt.Sprintf("age = $%d, ", argIndex)
		args = append(args, u.Age)
		argIndex++
	}

	if u.Email != "" {
		query += fmt.Sprintf("email = $%d, ", argIndex)
		args = append(args, u.Email)
		argIndex++
	}

	if u.Password != "" {
		query += fmt.Sprintf("password = $%d, ", argIndex)
		args = append(args, u.Password)
		argIndex++
	}

	// Видаляємо останню кому
	query = query[:len(query)-2]
	query += fmt.Sprintf(" WHERE id = $%d", argIndex)
	args = append(args, id)

	// Якщо немає полів для обновлення
	if len(args) == 1 {
		log.WithFields(log.Fields{
			"layer":  "storage",
			"action": "updateUser",
		}).Warn(fmt.Sprintf("No fields to update for user id %d", id))
		return nil
	}

	res, err := s.DB.Exec(query, args...)

	// перевіряцмо чи примінились нащі оновлення
	if err := ensureRowsAffected(res, id); err != nil {
		return err
	}

	if err != nil {
		log.WithFields(log.Fields{
			"layer":  "storage",
			"action": "updateUser",
		}).Error(err)
		//Check what i need return after logging errors
		return nil
	}

	log.WithFields(log.Fields{
		"layer":  "storage",
		"action": "updateUser",
	}).Info(fmt.Sprintf("Successfully updated user with id %d", id))

	return nil
}

func (s *Storage) DeleteUser(id int64) error {
	res, err := s.DB.Exec("DELETE FROM  users WHERE id = $1", id)

	if err != nil {
		return err
	}

	// перевіряцмо чи примінились нащі оновлення
	if err := ensureRowsAffected(res, id); err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"layer":  "storage",
		"action": "deleteUser",
	}).Info(fmt.Sprintf("Successfully delete user with id %d", id))

	return nil
}

func ensureRowsAffected(res sql.Result, id int64) error {

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New(fmt.Sprintf("User not found with id = %d", id))
	}

	return nil
}
